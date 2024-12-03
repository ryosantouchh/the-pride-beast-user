package repository

import (
	"fmt"

	"github.com/ryosantouchh/the-pride-beast-user/internal/core/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(newUser *entity.User) error {
	result := r.db.Create(newUser)

	if result.Error != nil {
		return fmt.Errorf("UserRepository[Create] failed to create new user : %w", result.Error)
	}
	return nil
}

func (r *UserRepositoryImpl) Get() (*[]entity.User, error) {
	var users []entity.User
	result := r.db.Find(&users)

	if result.Error != nil {
		return nil, fmt.Errorf("UserRepository[Get] failed to find users : %w", result.Error)
	}
	return &users, nil
}

func (r *UserRepositoryImpl) GetById(userId int) (*entity.User, error) {
	var users entity.User
	result := r.db.Where("id = ?", userId).First(&users)

	if result.Error != nil {
		return nil, fmt.Errorf("UserRepository[GetById] failed to find user by id : %w", result.Error)
	}
	return &users, nil
}

func (r *UserRepositoryImpl) Update(userId int, data interface{}) error {
	result := r.db.Model(entity.User{}).Where("id = ?", userId).Updates(data)

	if result.Error != nil {
		return fmt.Errorf("UserRepository[Update] failed to update user by id : %w", result.Error)
	}
	return nil
}

func (r *UserRepositoryImpl) SoftDelete(userId int, data interface{}) error {
	result := r.db.Model(entity.User{}).Where("id = ?", userId).Updates(data)

	if result.Error != nil {
		return fmt.Errorf("UserRepository[Update] failed to update user by id : %w", result.Error)
	}
	return nil
}
