package infrastructure

import (
	"github.com/garyburd/redigo/redis"
)

// CacheHandler struct.
type CacheHandler struct {
	// Cache is global redis connect
	Conn *redis.Conn
}

// NewCacheHandler returns new cacheHandler.
// repository: https://github.com/garyburd/redigo/redis
func NewCacheHandler() *CacheHandler {
	host := GetConfigString("redis.host")
	port := GetConfigString("redis.port")
	pass := GetConfigString("redis.pass")

	var err error
	options := redis.DialPassword(pass)
	c, err := redis.Dial("tcp", host+":"+port, options)

	if err != nil {
		panic(err)
	}
	return &CacheHandler{Conn: &c}
}
