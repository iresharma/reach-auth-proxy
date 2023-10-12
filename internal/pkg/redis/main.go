package redis

import (
	database "awesomeProject/internal/pkg/database"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var Rdb *redis.Client

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

func SessionExists(authId string) bool {
	ctx := context.Background()
	_, err := Rdb.Get(ctx, authId).Result()
	if err != nil {
		return false
	}
	return true
}

func DeleteSessionCache(authId string) {
	ctx := context.Background()
	sessId := Rdb.Get(ctx, authId)
	_ = Rdb.Del(ctx, authId)
	_ = Rdb.Del(ctx, sessId.String())
	_ = Rdb.Del(ctx, sessId.String()+":"+authId)
}

func AddSessionToCache(authId string, sessionId string, perm string) {
	ctx := context.Background()
	_ = Rdb.Set(ctx, authId, sessionId, 5*86400*time.Second)
	_ = Rdb.Set(ctx, sessionId, authId, 5*86400*time.Second)
	_ = Rdb.Set(ctx, sessionId+":"+authId, perm, 5*86400*time.Second)
}

func FetchSessionCache(sessionId string) (*map[string]string, *string) {
	ctx := context.Background()
	authId, err := Rdb.Get(ctx, sessionId).Result()
	if err != nil {
		fmt.Println(err)
		return database.FetchSessionDB(sessionId)
	}
	perm, err := Rdb.Get(ctx, sessionId+":"+authId).Result()
	data := map[string]string{
		"sessionId": sessionId,
		"authId":    authId,
		"perm":      perm,
	}
	return &data, nil
}
