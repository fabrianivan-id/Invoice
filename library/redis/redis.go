package redis

import (
	"context"

	"esb-test/library/logger"

	"github.com/redis/go-redis/v9"
)

func Init(ctx context.Context, addr, password string) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	}

	redisClient := redis.NewClient(opts)
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		logger.GetLogger(ctx).Errorf("init redis fail: ", err)
		return nil, err
	}
	//redisClient.AddHook(nrredis.NewHook(opts))
	return redisClient, nil
}
