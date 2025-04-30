package usecase

import "backend/domain"

type UserUsecase interface {
	SignUp(u *domain.User) error
}
