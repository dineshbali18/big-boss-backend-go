package cache

import (
	"big-boss-7/config"
	"big-boss-7/domain"
	"context"
	"fmt"
	"strings"

	redis "github.com/redis/go-redis/v9"
)

type redisCacheService struct {
	redisClient redis.UniversalClient
}

var ctx = context.Background()

func NewRedisCacheService(redisClient redis.UniversalClient) domain.CacheService {
	return &redisCacheService{
		redisClient: redisClient,
	}
}

// InitRedisCacheService initializes redis cache service
func InitRedisCacheService() redis.UniversalClient {
	rdbConfig, err := config.GetRedisConfig()
	if err != nil {
		panic(err)
	}
	var redisClient redis.UniversalClient
	if rdbConfig.RedisClusterModeEnabled {
		addr := strings.Split(rdbConfig.RedisURL, ",")
		redisClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:         addr,
			Password:      rdbConfig.RedisPassword,
			RouteRandomly: true,
			DialTimeout:   config.DialTimeout,
			ReadTimeout:   config.ReadTimeout,
			WriteTimeout:  config.WriteTimeout,
		})
	} else {
		fmt.Println("ddddddddddd", rdbConfig.RedisURL)
		redisClient = redis.NewClient(&redis.Options{
			Addr:         rdbConfig.RedisURL,
			Password:     rdbConfig.RedisPassword,
			DialTimeout:  config.DialTimeout,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
		})
	}

	res, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("WWWWWWWWWW", err)
	} else {
		fmt.Println("RRRRRRRRR", res)
	}

	return redisClient
}

func (rdb *redisCacheService) CheckRedisConnection() (result string, err error) {
	result, err = rdb.redisClient.Ping(ctx).Result()
	return
}
