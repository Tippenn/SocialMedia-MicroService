package wizard

import (
	"github.com/Tippenn/Social-Media-Project/config"
	sessionRepo "github.com/Tippenn/Social-Media-Project/domains/JWTVerification/repository"
	sessionUc "github.com/Tippenn/Social-Media-Project/domains/JWTVerification/usecase"
	"github.com/Tippenn/Social-Media-Project/infrastructure/cache"
)

var (
	Config = config.GetConfig()

	//infra
	Redis = cache.NewRedisDatabase(Config)

	//repo
	SessionRepo = sessionRepo.NewServerJWTRepository(Redis)

	//usecase
	SessionUsecase = sessionUc.NewServerJWTUseCase(SessionRepo)
)
