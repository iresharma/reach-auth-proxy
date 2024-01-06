package redis

import (
	database "awesomeProject/internal/pkg/database"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

var Rdb *redis.Client

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS"),
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
	_ = Rdb.Set(ctx, authId, sessionId, 1*86400*time.Second)
	_ = Rdb.Set(ctx, sessionId, authId, 1*86400*time.Second)
	_ = Rdb.Set(ctx, sessionId+":"+authId, perm, 1*86400*time.Second)
}

func FetchSessionCache(sessionId string) (*map[string]string, *string) {
	ctx := context.Background()
	authId, err := Rdb.Get(ctx, sessionId).Result()
	if err != nil {
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

func GetItemFromRedis(key string) interface{} {
	ctx := context.Background()
	val, err := Rdb.Get(ctx, key).Result()
	var data map[string]interface{}
	json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil
	}
	return data
}

func AddItemToRedis(key string, val interface{}) {
	ctx := context.Background()
	valMarshal, _ := json.Marshal(val)
	out := Rdb.Set(ctx, key, valMarshal, 1*86400*time.Second)
	log.Println(out.Result())
}

func DeleteFromRedis(key string) {
	ctx := context.Background()
	_ = Rdb.Del(ctx, key)
}

func DeleteAllKeysPrefix(prefix string) {
	ctx := context.Background()
	var cursor uint64
	for {
		var keys []string
		var err error

		keys, cursor, err = Rdb.Scan(ctx, cursor, prefix+"*", 10).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("deleting key: ", key)
			Rdb.Del(ctx, key)
		}

		if cursor == 0 {
			break
		}
	}
}
