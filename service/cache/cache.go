package cache

import "github.com/gomodule/redigo/redis"

type Broker interface {
	C() *redis.Pool
}
