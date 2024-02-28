package redis

import "github.com/go-redis/redis/v7"

// 添加列表
func List(key string, score float64, val []byte) error {
	return RDB.ZAdd(key, &redis.Z{
		Score:  score,
		Member: val,
	}).Err()
}

// 获取redis列表
func GetList(key string, start int64, stop int64) ([]string, error) {
	return RDB.ZRevRange(key, start, stop).Result()
}

// 检查是否存在键,存在返回true
func RedisKeyExists(key string) bool {
	ok, _ := RDB.Exists(key).Result()
	return ok == 1
}
