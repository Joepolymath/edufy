package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var (
	ErrInternal    = errors.New("Redis Driver Failure")
	ErrKeyNotFound = errors.New("Key Not Found")
)

type IRedisDriver interface {
	GetValueString(ctx context.Context, key string) (*string, error)
	SetValue(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	DeleteValue(ctx context.Context, key string) error
}

type RedisDriver struct {
	client *redis.Client
}

func NewRedisDriver(client *redis.Client) IRedisDriver {
	driver := &RedisDriver{
		client,
	}
	return driver
}

// Connect handles connection to data source server implementation
func RedisConnection() *redis.Client {
	ctx := context.Background()
	var rdb *redis.Client

	if strings.ToLower(cfg.ENVIRON) != "development" {
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.RedisAddress,
			Password: cfg.RedisPassword, // no password set
			Username: cfg.RedisUsername,
			DB:       0, // use default DB
		})
	} else {

		rdb = redis.NewClient(&redis.Options{
			Addr: cfg.RedisAddress,
			DB:   0, // use default DB
		})
	}

	st := rdb.Ping(ctx)
	if err := st.Err(); err != nil {
		log.Fatal("connection to redis server failed::: ", err)
		return nil
	}
	log.Print("successfully connected to redis server as user:: ", cfg.RedisUsername)
	return rdb
}

func (driver *RedisDriver) GetValueString(ctx context.Context,
	key string) (*string, error) {
	// Get the value for the key
	result := driver.client.Get(ctx, key)
	if result.Err() == redis.Nil {
		return nil, ErrKeyNotFound
	} else if result.Err() != nil {
		// Some other error occurred
		fmt.Println("Error:", result.Err())
		return nil, ErrInternal
	} else {
		// Key found, retrieve the value
		value := result.Val()
		return &value, nil

	}

}

func (driver *RedisDriver) SetValue(ctx context.Context, key string,
	value interface{}, ttl time.Duration) error {

	// set the value in memory
	res := driver.client.Set(ctx, key, value, ttl)
	if res.Err() != nil {
		fmt.Println("Error:", res.Err())
		return ErrInternal
	}

	return nil
}

func (driver *RedisDriver) DeleteValue(ctx context.Context, key string) error {
	res := driver.client.Del(ctx, key)
	if res.Err() != nil {
		fmt.Println("Error:", res.Err())
		return ErrInternal
	}
	return nil
}
