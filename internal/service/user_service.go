package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/nalendro16/learn-vibe-code/internal/dto"
	"github.com/nalendro16/learn-vibe-code/internal/model"
	"github.com/nalendro16/learn-vibe-code/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email is already registered")
)

// UserService defines the business logic operations for users
type UserService interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, req dto.RegisterRequest) error {
	// 1. Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return fmt.Errorf("failed to check existing email: %w", err)
	}
	if existingUser != nil {
		return ErrEmailAlreadyRegistered
	}

	// 2. Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// 3. Create user entity
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// 4. Save to repository
	if err := s.userRepo.Create(ctx, user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}
