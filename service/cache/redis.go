package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/varmamsp/cello/model"
)

type supplier struct {
	pool *redis.Pool
}

func NewBroker(config *model.Config) (Broker, error) {
	pool := &redis.Pool{
		MaxIdle: config.Redis.MaxIdleConn,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.Redis.Address)
		},
	}

	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return nil, err
	}
	return &supplier{pool: pool}, nil
}

func (splr *supplier) C() *redis.Pool {
	return splr.pool
}
