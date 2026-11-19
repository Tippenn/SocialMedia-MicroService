package usecase

import "github.com/Tippenn/Social-Media-Project/domain/users"

type userUsecase struct{
	repo users.UserRepository
}

func NewUserUsecase(repo users.UserRepository) users.UsersUsecase{
	return userUsecase{repo: repo}
}
