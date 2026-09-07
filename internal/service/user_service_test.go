package service

import (
	"context"
	"errors"
	"testing"

	"github.com/nalendro16/learn-vibe-code/internal/dto"
	"github.com/nalendro16/learn-vibe-code/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepository struct {
	findByEmailFunc func(ctx context.Context, email string) (*model.User, error)
	createFunc      func(ctx context.Context, user *model.User) error
}

func (m *mockUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	if m.findByEmailFunc != nil {
		return m.findByEmailFunc(ctx, email)
	}
	return nil, nil
}

func (m *mockUserRepository) Create(ctx context.Context, user *model.User) error {
	if m.createFunc != nil {
		return m.createFunc(ctx, user)
	}
	return nil
}

func TestUserService_Register(t *testing.T) {
	t.Run("Success Register", func(t *testing.T) {
		var savedUser *model.User
		repo := &mockUserRepository{
			findByEmailFunc: func(ctx context.Context, email string) (*model.User, error) {
				return nil, nil
			},
			createFunc: func(ctx context.Context, user *model.User) error {
				savedUser = user
				return nil
			},
		}

		svc := NewUserService(repo)
		req := dto.RegisterRequest{
			Name:     "darsam",
			Email:    "darsam@gmail.com",
			Password: "plainPassword123",
		}

		err := svc.Register(context.Background(), req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if savedUser == nil {
			t.Fatal("expected user to be saved")
		}

		if savedUser.Name != req.Name || savedUser.Email != req.Email {
			t.Errorf("saved user name/email mismatch")
		}

		// Verify password was hashed with bcrypt
		if err := bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(req.Password)); err != nil {
			t.Errorf("password was not properly hashed with bcrypt: %v", err)
		}
	})

	t.Run("Email Already Exists", func(t *testing.T) {
		repo := &mockUserRepository{
			findByEmailFunc: func(ctx context.Context, email string) (*model.User, error) {
				return &model.User{ID: "existing-id", Email: email}, nil
			},
		}

		svc := NewUserService(repo)
		req := dto.RegisterRequest{
			Name:     "darsam",
			Email:    "darsam@gmail.com",
			Password: "plainPassword123",
		}

		err := svc.Register(context.Background(), req)
		if !errors.Is(err, ErrEmailAlreadyRegistered) {
			t.Errorf("expected ErrEmailAlreadyRegistered, got %v", err)
		}
	})
}
