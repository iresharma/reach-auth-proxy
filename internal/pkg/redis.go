package pkg

import (
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
	if err != redis.Nil {
		return false
	} else if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func DeleteSession(authId string) {
	ctx := context.Background()
	sessId := Rdb.Get(ctx, authId)
	_ = Rdb.Del(ctx, authId)
	_ = Rdb.Del(ctx, sessId.String())
}

func AddSessionToCache(authId string, sessionId string) {
	ctx := context.Background()
	_ = Rdb.Set(ctx, authId, sessionId, 5*86400*time.Second)
	_ = Rdb.Set(ctx, sessionId, authId, 5*86400*time.Second)
}
