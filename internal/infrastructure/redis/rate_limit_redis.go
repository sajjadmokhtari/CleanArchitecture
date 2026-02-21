package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiterRedis struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRateLimiterRedis(client *redis.Client) *RateLimiterRedis {
	return &RateLimiterRedis{
		Client: client,
		Ctx:    context.Background(),
	}
}

func (r *RateLimiterRedis) CanSendOTP(phone string) (bool, error) {
	key := "otp:last:" + phone
	lastTimeStr, err := r.Client.Get(r.Ctx, key).Result()

	if err == redis.Nil {
		return true, nil
	}
	if err != nil {
		return true, err
	}

	lastTime, _ := strconv.ParseInt(lastTimeStr, 10, 64)
	return time.Now().Unix()-lastTime >= 30, nil
}

func (r *RateLimiterRedis) MarkOTPSent(phone string) error {
	key := "otp:last:" + phone
	return r.Client.Set(r.Ctx, key, time.Now().Unix(), time.Hour).Err()
}

func (r *RateLimiterRedis) OTPRequestCount(phone string) (int, error) {
	key := "otp:count:" + phone
	countStr, err := r.Client.Get(r.Ctx, key).Result()

	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(countStr)
}

func (r *RateLimiterRedis) IncrementOTPRequest(phone string) error {
	key := "otp:count:" + phone
	if err := r.Client.Incr(r.Ctx, key).Err(); err != nil {
		return err
	}
	return r.Client.Expire(r.Ctx, key, time.Hour).Err()
}

func (r *RateLimiterRedis) GetFailedAttempts(phone string) (int, error) {
	key := "otp:fail:" + phone
	countStr, err := r.Client.Get(r.Ctx, key).Result()

	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(countStr)
}

func (r *RateLimiterRedis) IncrementFailedAttempts(phone string) error {
	key := "otp:fail:" + phone
	if err := r.Client.Incr(r.Ctx, key).Err(); err != nil {
		return err
	}
	return r.Client.Expire(r.Ctx, key, time.Hour).Err()
}
