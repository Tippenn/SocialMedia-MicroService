package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Tippenn/Social-Media-Project/domains/JWTVerification"
	"github.com/Tippenn/Social-Media-Project/domains/JWTVerification/entities"
	"github.com/Tippenn/Social-Media-Project/infrastructure/cache"

	"github.com/redis/go-redis/v9"
)

type serverJWTRepository struct {
	cache cache.Cache
}

func NewServerJWTRepository(cache cache.Cache) JWTVerification.ServerJWTRepository {
	return serverJWTRepository{cache: cache}
}

// FindServerJWTbyUserID implements [JWTVerification.ServerJWTRepository].
func (s serverJWTRepository) FindServerJWTbyUserID(ctx context.Context, userId string) (*entities.ServerJWT, error) {
	var user *entities.ServerJWT
	result, err := s.cache.GetInstance().Get(ctx, "UserId:"+userId).Result()
	if errors.Is(err, redis.Nil) {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &user)
	return user, nil
}
