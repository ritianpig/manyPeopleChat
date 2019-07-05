package redisdb

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// 定义全局变量Pool redis连接池
var Pool *redis.Pool

func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
