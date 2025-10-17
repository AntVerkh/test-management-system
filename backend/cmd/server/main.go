package main

import (
	"log"

	"github.com/AntVerkh/test-management-system/internal/config"
	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/AntVerkh/test-management-system/internal/handler"
	"github.com/AntVerkh/test-management-system/internal/middleware"
	"github.com/AntVerkh/test-management-system/internal/repository"
	"github.com/AntVerkh/test-management-system/internal/service"
	"github.com/AntVerkh/test-management-system/pkg/auth"
	"github.com/AntVerkh/test-management-system/pkg/database"
	"github.com/AntVerkh/test-management-system/pkg/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize storage (remove unused variable warning by using it)
	fileStorage := storage.NewLocalFileStorage(cfg.FileStoragePath)
	_ = fileStorage // Use the variable to avoid "unused variable" error

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	testPlanRepo := repository.NewTestPlanRepository(db)
	testCaseRepo := repository.NewTestCaseRepository(db)
	checklistRepo := repository.NewChecklistRepository(db)
	testStrategyRepo := repository.NewTestStrategyRepository(db)
	testRunRepo := repository.NewTestRunRepository(db)

	// Initialize services
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	authService := service.NewAuthService(userRepo, jwtService)
	testPlanService := service.NewTestPlanService(testPlanRepo)
	testCaseService := service.NewTestCaseService(testCaseRepo)
	userService := service.NewUserService(userRepo)
	exporter := domain.NewMarkdownExporter()
	exportService := service.NewExportService(
		testPlanRepo,
		testCaseRepo,
		checklistRepo,
		testStrategyRepo,
		testRunRepo,
		exporter,
	)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService, userService)
	testPlanHandler := handler.NewTestPlanHandler(testPlanService)
	testCaseHandler := handler.NewTestCaseHandler(testCaseService)
	exportHandler := handler.NewExportHandler(exportService)

	// Setup router
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/auth/login", authHandler.Login)
		public.POST("/auth/register", authHandler.Register)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(authService))
	{
		// User routes
		protected.GET("/profile", authHandler.GetProfile)

		// Test Plans
		protected.GET("/test-plans", testPlanHandler.ListTestPlans)
		protected.POST("/test-plans", middleware.RoleMiddleware(domain.RoleAdmin, domain.RoleUser), testPlanHandler.CreateTestPlan)
		protected.GET("/test-plans/:id", testPlanHandler.GetTestPlan)
		protected.PUT("/test-plans/:id", middleware.RoleMiddleware(domain.RoleAdmin, domain.RoleUser), testPlanHandler.UpdateTestPlan)
		protected.POST("/test-plans/:id/test-cases", middleware.RoleMiddleware(domain.RoleAdmin, domain.RoleUser), testPlanHandler.AddTestCase)

		// Test Cases
		protected.GET("/test-cases", testCaseHandler.ListTestCases)
		protected.POST("/test-cases", middleware.RoleMiddleware(domain.RoleAdmin, domain.RoleUser), testCaseHandler.CreateTestCase)
		protected.GET("/test-cases/:id", testCaseHandler.GetTestCase)
		protected.PUT("/test-cases/:id", middleware.RoleMiddleware(domain.RoleAdmin, domain.RoleUser), testCaseHandler.UpdateTestCase)

		// Export routes
		protected.POST("/export", exportHandler.Export)
		protected.GET("/test-plans/:id/export", exportHandler.ExportTestPlan)
		protected.GET("/test-cases/:id/export", exportHandler.ExportTestCase)
		protected.GET("/checklists/:id/export", exportHandler.ExportChecklist)
		protected.GET("/test-strategies/:id/export", exportHandler.ExportTestStrategy)
		protected.GET("/test-runs/:id/export", exportHandler.ExportTestRun)

		// Admin only routes
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware(domain.RoleAdmin))
		{
			admin.GET("/users", authHandler.ListUsers)
			admin.PUT("/users/:id/role", authHandler.UpdateUserRole)
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
