package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const urlListKey = "urlListKey"

type RedisClient struct {
	ctx    context.Context
	client *redis.Client
}

func (r *RedisClient) Insert(url string) (id int64, err error) {
	res, err := r.client.RPush(r.ctx, urlListKey, url).Result()
	if err != nil {
		return 0, err
	}

	return res - 1, nil
}

func (r *RedisClient) Retrieve(id int64) (url string, err error) {
	res, err := r.client.LIndex(r.ctx, urlListKey, id).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func NewRedisClient(ctx context.Context, client *redis.Client) *RedisClient {
	return &RedisClient{
		ctx:    ctx,
		client: client,
	}
}
