package infrastructure

import (
	"article/domain"
	"article/domain/repository"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.IUserRepository {
	return &Repository{db}
}

func (r *Repository) Save(user domain.User) (domain.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) FindByEmail(email string) (domain.User, error) {
	var user domain.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
