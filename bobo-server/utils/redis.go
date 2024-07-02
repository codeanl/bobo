package utils

import (
	"bobo-server/config"
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"time"
)

var Redis = new(_redis)
var rdb *redis.Client
var ctx = context.Background()

func InitRedis() *redis.Client {
	redisCfg := config.Cfg.Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	// 测试连接状况
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic("Reids 连接失败: ", err)
	}
	log.Println("Redis 连接成功 ")
	return rdb
}

type _redis struct{}

// GetVal 获取值
func (*_redis) GetVal(key string) string {
	return rdb.Get(ctx, key).Val()
}

// Set 设置key value 过期时间
func (*_redis) Set(key string, value interface{}, expiration time.Duration) {
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		Logger.Error("utils/redis.go -> Set: ", zap.Error(err))
		panic(err)
	}
}

// 获取[有序集合]中, 成员的分数值
func (*_redis) ZScore(key, member string) int {
	return int(rdb.ZScore(ctx, key, member).Val())
}

// [有序集合]中 key 中指定字段的整数值加上增量 incr
func (*_redis) ZincrBy(key, member string, incr float64) {
	err := rdb.ZIncrBy(ctx, key, incr, member).Err()
	if err != nil {
		Logger.Error("utils/redis.go ->ZincrBy: ", zap.Error(err))
		panic(err)
	}
}
