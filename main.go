package main

import (
	"github.com/go-redis/redis"
	"route"
	"setUpRedis"
)



func main() {
	var rdb *redis.Client
	setUpRedis.InitClient(rdb)
	r:=route.SetUpRount()
	r.Run(":8080")

}




