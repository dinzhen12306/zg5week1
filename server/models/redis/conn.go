package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"week1/server/config"
)

var RDB *redis.Client

// 初始化redis连接
func RedisConn(m *config.Redis) (*redis.Client, error) {
	RDB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", m.Host, m.Port),
	})
	if _, err := RDB.Ping().Result(); err != nil {
		return nil, err
	}
	return RDB, nil
}
