package cache

import (
	"big-boss-7/config"
	"big-boss-7/domain"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "time/tzdata"

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
			PoolSize:      15,
			PoolTimeout:   1500 * time.Second,
			DialTimeout:   config.DialTimeout,
			ReadTimeout:   config.ReadTimeout,
			WriteTimeout:  config.WriteTimeout,
		})
	} else {
		fmt.Println("ddddddddddd", rdbConfig.RedisURL)
		redisClient = redis.NewClient(&redis.Options{
			Addr:         rdbConfig.RedisURL,
			Password:     rdbConfig.RedisPassword,
			PoolSize:     15,
			PoolTimeout:  1500 * time.Second,
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

func (rdb *redisCacheService) GetAllContestants() (contestants []domain.Contestants, err error) {
	cacheKey := "all:contestants"
	response, redisErr := rdb.redisClient.Get(context.Background(), cacheKey).Result()
	if redisErr != nil {
		fmt.Println("Cannot get data in get all contestants")
		return contestants, redisErr
	}

	err = json.Unmarshal([]byte(response), &contestants)
	if err != nil {
		return
	}
	return

}

func (rdb *redisCacheService) SaveAllContestants(contestants []domain.Contestants) (err error) {
	cacheKey := "all:contestants"

	keyExists := rdb.redisClient.Exists(context.Background(), cacheKey)
	if keyExists.Val() == 0 {
		marshaledData, marshalErr := json.Marshal(contestants)
		if marshalErr != nil {
			return marshalErr
		}
		err = rdb.redisClient.Set(context.Background(), cacheKey, string(marshaledData), config.ContestantsDataTTL).Err()
		if err != nil {
			fmt.Println("Error in saving data in cache")
			return
		}
	}
	return
}

func (rdb *redisCacheService) GetNominatedContestants() (contestants []domain.Contestants, err error) {
	cacheKey := "nominated:contestants"

	loc, err1 := time.LoadLocation("Asia/Kolkata")
	if err1 != nil {
		fmt.Println("Error loading location:", err1)
	}

	// Get the current day as a string
	day := time.Now().In(loc).Format("2006-01-02") // Format it as needed

	fmt.Println("TIME", time.Now())
	fmt.Println("TIME", time.Now().Format("2006-01-02"))

	// Append the day to the cacheKey
	cacheKey += ":" + day

	response, redisErr := rdb.redisClient.Get(context.Background(), cacheKey).Result()
	if redisErr != nil {
		return contestants, redisErr
	}

	err = json.Unmarshal([]byte(response), &contestants)
	if err != nil {
		fmt.Println("Failed to unmarshal in get nominated contestants")
		return
	}
	return
}

func (rdb *redisCacheService) SaveNominatedContestants(contestants []domain.Contestants) (err error) {
	cacheKey := "nominated:contestants"
	loc, err1 := time.LoadLocation("Asia/Kolkata")
	if err1 != nil {
		fmt.Println("Error loading location:", err1)
	}

	// Get the current day as a string
	day := time.Now().In(loc).Format("2006-01-02") // Format it as needed

	// Append the day to the cacheKey
	cacheKey += ":" + day

	keyExists := rdb.redisClient.Exists(context.Background(), cacheKey)
	if keyExists.Val() == 0 {
		marshaledData, marshalErr := json.Marshal(contestants)
		if marshalErr != nil {
			return marshalErr
		}
		err = rdb.redisClient.Set(context.Background(), cacheKey, string(marshaledData), config.NominatedContestantsTTL).Err()
		if err != nil {
			fmt.Println("Error in saving data in cache")
			return
		}
	}
	return
}

func (rdb *redisCacheService) GetPercentagesResults() (voteData domain.VotesPercentages, err error) {
	cacheKey := "nominated:contestants:votes"
	loc, err1 := time.LoadLocation("Asia/Kolkata")
	if err1 != nil {
		fmt.Println("Error loading location:", err1)
	}

	// Get the current day as a string
	day := time.Now().In(loc).Format("2006-01-02") // Format it as needed

	// Append the day to the cacheKey
	cacheKey += ":" + day
	response, redisErr := rdb.redisClient.Get(context.Background(), cacheKey).Result()
	if redisErr != nil {
		return voteData, redisErr
	}
	err = json.Unmarshal([]byte(response), &voteData)
	if err != nil {
		return
	}
	return
}

func (rdb *redisCacheService) SavePercentagesResults(VotesData domain.VotesPercentages) (err error) {
	cacheKey := "nominated:contestants:votes"

	loc, err1 := time.LoadLocation("Asia/Kolkata")
	if err1 != nil {
		fmt.Println("Error loading location:", err1)
	}

	// Get the current day as a string
	day := time.Now().In(loc).Format("2006-01-02") // Format it as needed

	// Append the day to the cacheKey
	cacheKey += ":" + day

	keyExists := rdb.redisClient.Exists(context.Background(), cacheKey)
	if keyExists.Val() == 0 {
		marshaledData, marshalErr := json.Marshal(VotesData)
		if marshalErr != nil {
			return marshalErr
		}
		err = rdb.redisClient.Set(context.Background(), cacheKey, string(marshaledData), config.CheckVotingPercentagesTTL).Err()
		if err != nil {
			fmt.Println("Error in saving data in cache")
			return
		}
	}
	return
}
