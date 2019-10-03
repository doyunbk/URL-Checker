package model

import (
	"errors"
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

var NoUrlFound = errors.New("no url found")

// Initialize the pool of Redis database
// for creating and configuring a connection
func InitPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			redisURL := os.Getenv("REDIS_URL")
			if len(redisURL) == 0 {
				redisURL = "localhost:6379"
			}
			c, err := redis.Dial("tcp", redisURL)
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

	fmt.Printf("PONG %s : Redis Connected \n", s)

	return nil
}

// Get Redis pool
func GetPool() *redis.Pool {
	if pool == nil {
		pool = InitPool()
	}
	return pool
}

// GetURL is getting the values of URL from Redis database by searching keys
func GetURL(url string) (*URL, error) {
	conn := GetPool().Get()

	defer conn.Close()

	// Lookup all fields and values of urls in the database
	urls, err := redis.Values(conn.Do("HGETALL", ""+url))
	if err != nil {
		return nil, err
	} else if len(urls) == 0 {
		return nil, NoUrlFound
	}

	// Scan the struct of values of urls
	var lookupurl URL
	err = redis.ScanStruct(urls, &lookupurl)
	if err != nil {
		return nil, err
	}

	return &lookupurl, nil
}
