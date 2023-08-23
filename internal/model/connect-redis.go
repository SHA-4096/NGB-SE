package model

import (
	config "NGB-SE/internal/conf"
	"NGB-SE/internal/util"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisCtx    context.Context
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Config.DataBase.RedisAddr,
		Password: config.Config.DataBase.RedisPassword,
		DB:       config.Config.DataBase.RedisDB,
	})
	redisCtx = context.Background()
}

//
//Set a key-value pair
//
func SetKeyValuePair(key, value string) error {
	err := redisClient.Do(redisCtx, "set", key, value).Err()
	if err != nil {
		util.MakeInfoLog(fmt.Sprintf("Error at model.SetKeyValuePair,msg=%s", err.Error()))
		return err
	}
	return nil
}

//
//Set key expiration time in redis
//
func SetExpiration(key string, seconds int) error {
	err = redisClient.Do(redisCtx, "expire", key, seconds).Err()
	if err != nil {
		util.MakeInfoLog(fmt.Sprintf("Error at model.SetExpiration,msg=%s", err.Error()))
		return err
	}
	return nil
}

//
//Discard key expiration time in redis
//
func DiscardExpiration(key string) error {
	err = redisClient.Do(redisCtx, "persist", key).Err()
	if err != nil {
		util.MakeInfoLog(fmt.Sprintf("Error at model.DiscardExpiration,msg=%s", err.Error()))
		return err
	}
	return nil
}

//
//Get the value of a specific key in redis
//
func GetKeyValue(key string) (string, error) {
	val, err := redisClient.Get(redisCtx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil

}
