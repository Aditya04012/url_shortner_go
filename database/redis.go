package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis()*redis.Client{
	rdb:=redis.NewClient(&redis.Options{
		Addr: "192.168.1.105:6379",
		Password:"Aditya@04" ,
		DB: 0,
	})
	_,err:=rdb.Ping(context.Background()).Result()
	if err!=nil{
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	fmt.Println("Redis connected Sucess")
 return rdb
}
