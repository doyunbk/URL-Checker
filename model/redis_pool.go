package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

// Initialize the pool of Redis database
// for creating and configuring a connection
func InitPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// Verify connectivity for redis
func Ping(c redis.Conn) error {
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING %s : Redis Server Connected \n", s)

	return nil
}
