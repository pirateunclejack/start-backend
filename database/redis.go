package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func RedisDb() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr: redis_addr+":"+redis_port,
        Password: redis_password,
        DB: redis_database,
    })

    return rdb
}

func RedisSet(rdb *redis.Client, key string, value interface{}) error {
    var ctx = context.Background()
    err := rdb.Set(ctx, key, value, 0).Err()
    if err != nil {
        return err
    }
    return nil
}

func RedisGet(rdb *redis.Client, key string) (string, error){
    var ctx = context.Background()
    res := rdb.Get(ctx, key)
    err := res.Err()
    if err != nil {
        return "", err
    }
    return res.Val(), nil
}
