package rdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
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

func (r *Rdb) StoreMultiLevelHash(ctx context.Context, key string, data map[string]interface{}) {
	for k, v := range data {

		flatKey := r.flattenKey(key, k)
		if nestedMap, ok := v.(map[string]interface{}); ok {
			r.StoreMultiLevelHash(ctx, flatKey, nestedMap)
		} else {
			r.db.HSet(ctx, flatKey, k, v)
		}
	}
}

func (r *Rdb) flattenKey(prefix, key string) string {
	return fmt.Sprintf("%s.%s", prefix, key)
}

func (r *Rdb) RetrieveMultiLevelHash(ctx context.Context, key string) map[string]string {
	fields, err := r.db.HGetAll(ctx, key).Result()
	if err != nil {
		log.Fatalf("Failed to retrieve hash: %v", err)
	}
	return fields
}

func (r *Rdb) RetrieveNestedHash(ctx context.Context, prefix string) map[string]interface{} {
	nestedMap := make(map[string]interface{})

	iter := r.db.Scan(ctx, 0, prefix+"*", 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()
		relativeKey := key[len(prefix)+1:]
		parts := r.splitKey(relativeKey)

		currentMap := nestedMap
		for i, part := range parts {
			if i == len(parts)-1 {
				value, err := r.db.HGet(ctx, key, "value").Result()
				if errors.Is(err, redis.Nil) {
					nestedHash := r.RetrieveMultiLevelHash(ctx, key)
					currentMap[part] = nestedHash
				} else if err != nil {
					log.Fatalf("Failed to retrieve hash value: %v", err)
				} else {
					currentMap[part] = value
				}
			} else {
				if _, exists := currentMap[part]; !exists {
					currentMap[part] = make(map[string]interface{})
				}
				currentMap = currentMap[part].(map[string]interface{})
			}
		}
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("SCAN failed: %v", err)
	}

	return nestedMap
}

func (r *Rdb) splitKey(key string) []string {
	return strings.Split(key, ".")
}

func (r *Rdb) test() {
	data := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "John Doe",
			"age":  30,
			"address": map[string]interface{}{
				"street": "123 Main St",
				"city":   "Anytown",
			},
		},
		"preferences": map[string]interface{}{
			"notifications": true,
			"theme":         "dark",
		},
	}

	ctx := context.Background()

	for k, v := range data {
		r.StoreMultiLevelHash(ctx, k, v.(map[string]interface{}))
	}

	nestedHash := r.RetrieveNestedHash(ctx, "user")
	fmt.Println("Nested Hash:", nestedHash)
}
