package service

import (
	"context"
	"errors"

	"github.com/AntVerkh/test-management-system/internal/domain"
	"github.com/AntVerkh/test-management-system/internal/repository"
	"github.com/google/uuid"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context) ([]domain.User, error) {
	return s.userRepo.List(ctx)
}

func (s *userService) UpdateUserRole(ctx context.Context, userID uuid.UUID, role domain.UserRole) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Validate role
	switch role {
	case domain.RoleAdmin, domain.RoleUser, domain.RoleGuest:
		user.Role = role
	default:
		return errors.New("invalid role")
	}

	return s.userRepo.Update(ctx, user)
}
