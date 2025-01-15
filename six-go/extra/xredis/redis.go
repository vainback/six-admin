package xredis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx    = context.Background()
	client *redis.Client
)

func Client() *redis.Client {
	return do()
}

func do() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     cfg().addr(),
			Password: cfg().pass(),
			DB:       cfg().db(),
		})
	}

	if err := client.Ping(ctx).Err(); err != nil {
		log.Println("redis ping err:", err)
	}

	return client
}

func Ping() error {
	return do().Ping(ctx).Err()
}

func IsExpire(key string, seconds int64) bool {
	result := HasExpireSeconds(key)
	return result == -1 || result > float64(seconds)
}

func HasExpireSeconds(key string) float64 {
	result, err := do().TTL(ctx, key).Result()
	if err != nil {
		return -2
	}
	return result.Seconds()
}

func GetString(key string) string {
	result, err := do().Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func GetInt(key string) int64 {
	result, err := do().Get(ctx, key).Int64()
	if err != nil {
		return 0
	}
	return result
}

func Set(key string, value any) error {
	return do().Set(ctx, key, value, -1).Err()
}

// SetExp expiration 传入的值必须携带单位 例如：time.Second * 60 既有效期60秒  time.Hour * 3 既有效期3小时
func SetExp(key string, value any, expiration time.Duration) error {
	return do().Set(ctx, key, value, expiration).Err()
}

func Del(key string) error {
	return do().Del(ctx, key).Err()
}

func Expire(key string, exp time.Duration) error {
	return do().Expire(ctx, key, exp).Err()
}

func Push(key string, values ...any) error {
	return do().RPush(ctx, key, values...).Err()
}

func Pop(key string) (string, error) {
	return do().LPop(ctx, key).Result()
}
