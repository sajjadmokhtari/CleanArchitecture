package redis

import (
	"time"

	"CleanArchitecture/internal/repository"

	"github.com/redis/go-redis/v9"
)

type OtpRedisRepository struct {
	client *redis.Client
}

func NewOtpRedisRepository(client *redis.Client) repository.OTPRepository {
	return &OtpRedisRepository{
		client: client,
	}
}

func (r *OtpRedisRepository) Save(phone string, otp string, ttlSeconds int) error {
	return r.client.Set(
		Ctx,
		"otp:"+phone,
		otp,
		time.Duration(ttlSeconds)*time.Second,
	).Err()
}

func (r *OtpRedisRepository) Get(phone string) (string, error) {
	return r.client.Get(
		Ctx,
		"otp:"+phone,
	).Result()
}
