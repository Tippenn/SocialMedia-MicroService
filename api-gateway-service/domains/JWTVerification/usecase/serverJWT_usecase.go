package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Tippenn/Social-Media-Project/domains/JWTVerification"
	"github.com/Tippenn/Social-Media-Project/domains/JWTVerification/entities"
	"github.com/Tippenn/Social-Media-Project/domains/JWTVerification/models/request"
)

type serverJWTUseCase struct {
	repo JWTVerification.ServerJWTRepository
}

func NewServerJWTUseCase(repo JWTVerification.ServerJWTRepository) JWTVerification.ServerJWTUseCase {
	return serverJWTUseCase{repo: repo}
}

// CheckJWTServer implements [JWTVerification.ServerJWTUseCase].
func (s serverJWTUseCase) CheckJWTServer(ctx context.Context, requests *request.CheckJWTRequest) (*entities.ServerJWT, error) {
	jwtData, err := s.repo.FindServerJWTbyUserID(ctx, requests.UserId)
	if err != nil {
		return nil, err
	}

	if jwtData.Revoked {
		return nil, errors.New("token Revoked")
	}

	if time.Now().After(jwtData.ExpiresAt) {
		return nil, errors.New("token Expired")
	}

	return jwtData, nil
}
