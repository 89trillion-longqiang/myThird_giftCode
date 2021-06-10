package main

import (
	"giftCode/route"
	"giftCode/setUpRedis"
	"github.com/go-redis/redis"
)



func main() {
	var rdb *redis.Client
	setUpRedis.InitClient(rdb)
	r:=route.SetUpRount()
	r.Run(":8080")

}




