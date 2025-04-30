package usecase

import "backend/model"

type UserUsecase interface {
	CreateUser(u *model.User) error
}
