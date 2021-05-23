package redis

import (
	"context"
	"fmt"
	"strings"
	redis "github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	DB *redis.Client
}

func NewClient(ctx context.Context, host string, port string, db int) (*Client, error) {
	addr := strings.Join([]string{host, ":", port}, "")

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})

	pong, err := rdb.Ping(ctx).Result()

	log.WithFields(log.Fields{
		"result": fmt.Sprintf("Connection result to Redis: %v\n", pong),
	}).Debug("redis")

	if err != nil {
		return nil, err
	}

	return &Client{rdb}, nil
}

func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration int) error {
	return c.DB.Set(ctx, key, value, 0).Err()
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.DB.Get(ctx, key).Result()
}

func (c *Client) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	return c.DB.ZRevRangeWithScores(ctx, key, start, stop).Result()
}

func (c *Client) ZAdd(ctx context.Context, key string, members ...*Z) error {
	return c.DB.ZAdd(ctx, key, members...).Err()
}

type Z = redis.Z
