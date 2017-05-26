package data

import (
	"github.com/go-redis/redis"
	"../util"
)

var Cli = new(redis.Client)

//TODO panic when client is nil
func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     util.Redis.Host + ":" + util.Redis.Port,
		Password: "",
		DB:       0,
	})
	Cli = client
}
