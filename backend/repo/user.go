package repo

import (
	"backend/domain"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) SignUp(u *domain.User) error {
	return r.db.Create(u).Error
}
