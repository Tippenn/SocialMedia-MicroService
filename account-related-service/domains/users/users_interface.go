package users

import (
	"context"

	"github.com/Tippenn/Social-Media-Project/domains/users/entities"
	"github.com/Tippenn/Social-Media-Project/domains/users/models/request"
	"github.com/Tippenn/Social-Media-Project/domains/users/models/response"
)

type UserUseCase interface {
	Login(ctx context.Context, requests *request.LoginRequest) (*response.LoginResponse, error)
	SignUp(ctx context.Context, requests *request.SignUpRequest) (*response.SignUpResponse, error)
	FindUserbyEmail(email string) (*entities.User, error)
}

type UserRepository interface {
	FindUserbyEmail(email string) (*entities.User, error)
	FindUserbyID(Id string) (*entities.User, error)
	CreateUser(ctx context.Context, userId string, username string, name string, password string) error
	UpdateUser(username *entities.User) error
}
