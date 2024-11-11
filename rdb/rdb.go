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

func NewRdb(connString string) (*Rdb, error) {
	redisOnce.Do(func() {
		opt, err := redis.ParseURL(connString)
		if err != nil {
			panic(err)
		}

		db := redis.NewClient(opt)

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
