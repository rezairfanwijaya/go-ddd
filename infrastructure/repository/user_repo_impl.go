package infrastructure

import (
	"article/domain"
	"article/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.IUserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(user domain.User) (domain.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
