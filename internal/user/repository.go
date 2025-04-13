package user

import (
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(user *User) error
    FindByEmail(email string) (*User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) Create(user *User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
    var user User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
