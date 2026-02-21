package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // چون docker بالا آوردی
		Password: "",
		DB:       0,
	})
}
//در کل وصل شدن ردیس چیز خاصی نیست اینجا 