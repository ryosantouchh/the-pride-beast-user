package usecase

import "github.com/ryosantouchh/the-pride-beast-user/internal/core/service"

type FindUserByIdUsecase struct {
	userService service.UserService
}

func NewFindUserByIdUsecase(userService service.UserService) *FindUserByIdUsecase {
	return &FindUserByIdUsecase{userService: userService}
}

func (u *FindUserByIdUsecase) Execute() error {
	// err := u.userService
	return
}
