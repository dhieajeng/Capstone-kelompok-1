package cache

import (
	"fmt"
	"time"

	"github.com/bloomingbug/depublic/configs"
	"github.com/gomodule/redigo/redis"
)

func InitCache(config *configs.RedisConfig) *redis.Pool {
	rdb := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port))
		},
	}

	return rdb
}

type Cacheable interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) string
}

type cacheable struct {
	rdb *redis.Pool
}

func NewCacheable(rdb *redis.Pool) Cacheable {
	return &cacheable{
		rdb: rdb,
	}
}

func (c *cacheable) Set(key string, value interface{}, expiration time.Duration) error {
	conn := c.rdb.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value, "EX", int64(expiration.Seconds()))
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheable) Get(key string) string {
	conn := c.rdb.Get()
	defer conn.Close()

	val, err := redis.String(conn.Do("GET", key))
	if err == redis.ErrNil {
		return ""
	} else if err != nil {
		return ""
	}
	return val
}
