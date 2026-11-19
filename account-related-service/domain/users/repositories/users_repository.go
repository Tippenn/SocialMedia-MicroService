package repositories

import (
	"github.com/Tippenn/Social-Media-Project/domain/users"
	"github.com/Tippenn/Social-Media-Project/infrastructure/database"
)

type userUsecase struct {
	db users.UserRepository
}

func NewUserUsecase(db database.Database) users.UserRepository {
	return userUsecase{db: db}
}
