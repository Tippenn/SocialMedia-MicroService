package wizard

import (
	"github.com/Tippenn/Social-Media-Project/config"
	"github.com/Tippenn/Social-Media-Project/domains/users/handler/http"
	"github.com/Tippenn/Social-Media-Project/domains/users/repository"
	"github.com/Tippenn/Social-Media-Project/domains/users/usecase"
	"github.com/Tippenn/Social-Media-Project/infrastructure/cache"
	"github.com/Tippenn/Social-Media-Project/infrastructure/database"
)

var (
	Config = config.GetConfig()

	//infra
	PostgresDatabase = database.NewPostgresDatabase(Config)
	RedisCache       = cache.NewRedisDatabase(Config)

	//repo
	UserRepo = repository.NewUserRepository(PostgresDatabase, RedisCache)
	//usecase
	UserUsecase = usecase.NewUserUseCase(UserRepo)
	//http
	UserHttp = http.NewUserHttp(UserUsecase)
)
