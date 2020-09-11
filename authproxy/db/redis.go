package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/slovak-egov/einvoice/authproxy/user"
	"github.com/slovak-egov/einvoice/common"
)

var ctx = context.Background()

type AuthDB interface {
	Add(user *user.User)
	Remove(token string) error
	Exists(token string) bool
	GetUser(token string) *user.User
}

type redisDB struct {
	client *redis.Client
}

func NewAuthDB() AuthDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     common.GetRequiredEnvVariable("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default db
	})

	fmt.Println("ping", rdb.Ping(ctx).Val())

	return redisDB{rdb}
}

func (redisDB redisDB) Add(user *user.User) {
	redisDB.client.HSet(ctx, "user:"+user.Id, "token", user.Token)
	redisDB.client.HSet(ctx, "tokens", user.Token, user.Id)
}

func (redisDB redisDB) Remove(token string) error {
	id := redisDB.client.HGet(ctx, "tokens", token)
	if id.Val() == "" {
		return errors.New("not found")
	}
	redisDB.client.HDel(ctx, "tokens", token)
	redisDB.client.Del(ctx, "user:"+id.Val())
	return nil
}

func (redisDB redisDB) Exists(token string) bool {
	res := redisDB.client.HExists(ctx, "tokens", token)
	return res.Val()
}

func (redisDB redisDB) GetUser(token string) *user.User {
	id := redisDB.client.HGet(ctx, "tokens", token).Val()
	if id == "" {
		return nil
	}
	return &user.User{
		Token: token,
		Id:    id,
	}
}
