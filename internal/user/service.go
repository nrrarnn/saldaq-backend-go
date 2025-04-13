package user

import (
    "golang.org/x/crypto/bcrypt"
)

type UserService interface {
    Register(name, email, password string) (*User, error)
    Login(email, password string) (*User, error)
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

func (s *userService) Login(email, password string) (*User, error) {
    user, err := s.repo.FindByEmail(email)
    if err != nil {
        return nil, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, err
    }

    return user, nil
}
