package repository

import "gorm.io/gorm"

type Repository struct {
	User IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	UserRepo := NewUserRepository(db)
	return &Repository{
		User: UserRepo,
	}
}
