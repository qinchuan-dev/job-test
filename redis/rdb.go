package rdb

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

type Rdb struct {
	db *redis.Client
}

var (
	redisInstance *Rdb
	redisOnce     sync.Once
)

func NewRedis(connString string) (*Rdb, error) {
	redisOnce.Do(func() {
		db := redis.NewClient(&redis.Options{
			Addr: ":6379",
		})

		redisInstance = &Rdb{db}
	})

	return redisInstance, nil
}

func (r *Rdb) Close() {
	err := r.db.Close()
	if err != nil {
		return
	}
}
