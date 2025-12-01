package auth

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/mathieudottir/poker_tracker/backend/internal/models"
	"github.com/mathieudottir/poker_tracker/backend/internal/repository"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUserExists         = errors.New("user already exists")
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

// Register creates a new user
func (s *AuthService) Register(ctx context.Context, username, password, playerName string) (*models.User, error) {
	// Check if user already exists
	_, err := s.repo.GetUserByUsername(ctx, username)
	if err == nil {
		return nil, ErrUserExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		PlayerName:   playerName,
		WinaStatus:   "Aluminium",
		DevMode:      false,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// Login authenticates a user
func (s *AuthService) Login(ctx context.Context, username, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// DevLogin allows dev mode login without password (for "Login as Mathieu")
func (s *AuthService) DevLogin(ctx context.Context, username string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !user.DevMode {
		return nil, errors.New("dev mode not enabled for this user")
	}

	return user, nil
}

// UpdateUserSettings updates user settings
func (s *AuthService) UpdateUserSettings(ctx context.Context, userID int, playerName, winaStatus, hhDirectory string, devMode bool) error {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	user.PlayerName = playerName
	user.WinaStatus = winaStatus
	user.HHDirectory = hhDirectory
	user.DevMode = devMode

	return s.repo.UpdateUser(ctx, user)
}

// DeleteUserData deletes all data for a user
func (s *AuthService) DeleteUserData(ctx context.Context, userID int) error {
	return s.repo.DeleteUserData(ctx, userID)
}

// GetUser retrieves user by ID
func (s *AuthService) GetUser(ctx context.Context, userID int) (*models.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}
