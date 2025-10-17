.PHONY: build up down logs clean migrate test

# Build all services
build:
	docker-compose build

# Start all services
up:
	docker-compose up -d

# Stop all services
down:
	docker-compose down

# View logs
logs:
	docker-compose logs -f

# Clean up containers and volumes
clean:
	docker-compose down -v
	docker system prune -f

# Run database migrations
migrate:
	docker-compose exec backend ./migrate

# Run tests
test:
	docker-compose exec backend go test ./...

# Run backend in development mode
dev-backend:
	cd backend && go run cmd/server/main.go

# Run frontend in development mode
dev-frontend:
	cd frontend && npm run dev

# Generate swagger docs
swagger:
	cd backend && swag init -g cmd/server/main.go

# Deploy to production
deploy:
	git pull
	docker-compose -f docker-compose.prod.yml up -d --build