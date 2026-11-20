package JWTVerification

import (
	"context"

	"github.com/Tippenn/Social-Media-Project/domains/JWTVerification/entities"
	"github.com/Tippenn/Social-Media-Project/domains/JWTVerification/models/request"
)

type ServerJWTUseCase interface {
	CheckJWTServer(ctx context.Context, requests *request.CheckJWTRequest) (*entities.ServerJWT, error)
}

type ServerJWTRepository interface {
	FindServerJWTbyUserID(ctx context.Context, Id string) (*entities.ServerJWT, error)
}
