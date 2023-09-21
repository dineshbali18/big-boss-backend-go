package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// DialTimeout returns the dial timeout for the redis connection
var DialTimeout time.Duration = time.Millisecond * 1000

// ReadTimeout returns the read timeout for the redis connection
var ReadTimeout time.Duration = time.Millisecond * 1000

// WriteTimeout returns the write timeout for the redis connection
var WriteTimeout time.Duration = time.Millisecond * 1000

var CachePurgeEnabled bool = false

var ContestantsDataTTL time.Duration = 60 * 24 * 7 * time.Minute

var NominatedContestantsTTL time.Duration = 60 * 24 * time.Minute

var CheckVotingPercentagesTTL time.Duration = 20 * time.Minute

// RedisConfiguration holds the values required to connect to the redis instance
type RedisConfiguration struct {
	RedisClusterModeEnabled bool
	RedisHost               string
	RedisPort               string
	RedisPassword           string
	RedisURL                string
	RedisDB                 int
}

func GetRedisConfig() (config RedisConfiguration, err error) {
	// Set redis to cluster mode by default
	config.RedisClusterModeEnabled = true
	if viper.IsSet("REDIS_CLUSTER_MODE_ENABLED") {
		config.RedisClusterModeEnabled = viper.GetBool("REDIS_CLUSTER_MODE_ENABLED")
	}
	config.RedisHost = viper.GetString("REDIS_HOST")
	fmt.Println("co", config.RedisHost)
	config.RedisPort = viper.GetString("REDIS_PORT")
	config.RedisPassword = viper.GetString("REDIS_PASSWORD")
	config.RedisDB = viper.GetInt("REDIS_DB")
	CachePurgeEnabled = viper.GetBool("REDIS_CACHE_PURGE_ENABLED")
	contestantsDataTTL := viper.GetDuration("REDIS_CONTESTANTS_DATA_TTL")
	nominatedContestantsTTL := viper.GetDuration("REDIS_NOMINATED_CONTESTANTS_TTL")
	checkVotingPercentagesTTL := viper.GetDuration("REDIS_CHECK_PERCENTAGES_TTL")

	if contestantsDataTTL != 0 {
		ContestantsDataTTL = contestantsDataTTL * time.Second
	}

	if nominatedContestantsTTL != 0 {
		NominatedContestantsTTL = nominatedContestantsTTL * time.Second
	}

	if checkVotingPercentagesTTL != 0 {
		CheckVotingPercentagesTTL = checkVotingPercentagesTTL * time.Second
	}

	if len(config.RedisHost) == 0 {
		err = errors.New("invalid redis host")
		return
	}
	if len(config.RedisPort) == 0 {
		err = errors.New("invalid redis port")
		return
	}

	if config.RedisDB < 0 {
		err = errors.New("redis DB number is negative")
		return
	}

	DialTimeout = viper.GetDuration("REDIS_DIAL_TIMEOUT") * time.Millisecond
	ReadTimeout = viper.GetDuration("REDIS_READ_TIMEOUT") * time.Millisecond
	WriteTimeout = viper.GetDuration("REDIS_WRITE_TIMEOUT") * time.Millisecond

	config.RedisURL = fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)
	fmt.Println("RedisUrl", config.RedisURL)
	return
}
