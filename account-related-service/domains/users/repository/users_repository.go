package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Tippenn/Social-Media-Project/domains/users"
	"github.com/Tippenn/Social-Media-Project/domains/users/entities"
	"github.com/Tippenn/Social-Media-Project/infrastructure/cache"
	"github.com/Tippenn/Social-Media-Project/infrastructure/database"
	"github.com/redis/go-redis/v9"
)

type userRepository struct {
	db    database.Database
	cache cache.Cache
}

func NewUserRepository(db database.Database, cache cache.Cache) users.UserRepository {
	return userRepository{db: db, cache: cache}
}

// CreateUser implements [users.UserRepository].
func (repo userRepository) CreateUser(ctx context.Context, userId string, email string, name string, password string) error {
	user := &entities.User{
		Id:       userId,
		Email:    email,
		Name:     name,
		Password: password,
		Enabled:  1,
	}

	result := repo.db.GetInstance().Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindUserbyID implements [users.UserRepository].
func (repo userRepository) FindUserbyID(Id string) (*entities.User, error) {
	var user *entities.User
	result := repo.db.GetInstance().First(&user, "id=?", Id)

	if result.Error != nil {
		println("repo")
		return nil, result.Error
	}

	return user, nil
}

// FindUserbyUsername implements [users.UserRepository].
func (repo userRepository) FindUserbyEmail(email string) (*entities.User, error) {
	var user *entities.User
	result := repo.db.GetInstance().First(&user, "email=?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// UpdateUser implements [users.UserRepository].
func (repo userRepository) UpdateUser(username *entities.User) error {
	panic("unimplemented")
}

// CreateSession implements [users.UserRepository].
func (repo userRepository) CreateSession(ctx context.Context, accessToken string, userId string) (*entities.Session, error) {
	user := &entities.Session{
		Id:        accessToken,
		UserId:    userId,
		Revoked:   false,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	result := repo.cache.GetInstance().Set(ctx, "UserId:"+userId, data, time.Hour*1)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return user, nil
}

// DeleteSession implements [users.UserRepository].
func (repo userRepository) DeleteSession(ctx context.Context, userId string) error {
	result := repo.cache.GetInstance().Del(ctx, "UserId:"+userId)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

// FindSessionbyUserID implements [users.UserRepository].
func (repo userRepository) FindSessionbyUserID(ctx context.Context, Id string) (*entities.Session, error) {
	var user *entities.Session
	result, err := repo.cache.GetInstance().Get(ctx, "UserId:"+Id).Result()
	if errors.Is(err, redis.Nil) {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &user)
	return user, nil
}

// UpdateSession implements [users.UserRepository].
func (repo userRepository) UpdateSession(ctx context.Context, session *entities.Session) error {
	var user *entities.Session
	result, err := repo.cache.GetInstance().Get(ctx, "UserId:"+session.UserId).Result()
	if errors.Is(err, redis.Nil) {
		return err
	}
	err = json.Unmarshal([]byte(result), &user)

	user.Revoked = true

	data, _ := json.Marshal(user)
	repo.cache.GetInstance().Set(ctx, "UserId:"+session.UserId, data, redis.KeepTTL)

	return nil
}
