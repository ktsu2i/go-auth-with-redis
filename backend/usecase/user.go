package usecase

import "backend/model"

type UserUsecase interface {
	SignUp(u *model.User) error
}
