package user

import (
	"errors"
	"github.com/nrrarnn/saldaq-backend-go/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(name, email, password string) (*User, error)
	Login(email, password string) (map[string]interface{}, error) 
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(name, email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email, password string) (map[string]interface{}, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"token": token,
	}, nil
}
