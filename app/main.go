package main

import (
	"big-boss-7/config"
	"fmt"
	"log"

	cacheServices "big-boss-7/cacheservice"

	bbDelivery "big-boss-7/bb7/delivery/http"
	bbRepository "big-boss-7/bb7/repository/mysql"
	bbUsecase "big-boss-7/bb7/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	e *echo.Echo
)

func init() {
	//Initialize config
	config.InitializeConfig()
	e = echo.New()
}

func main() {
	//Load Database config from config.yml
	err := config.GetDatabaseConfig()
	if err != nil {
		log.Println(err.Error())
	}

	// Establish data base connection
	db, err := gorm.Open(mysql.Open(config.DatabaseConfig.DatabaseURL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Println(err.Error())
	}

	// Specifying DB Reader and Writer
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(config.DatabaseConfig.DatabaseWriteURL)},
		Replicas: []gorm.Dialector{mysql.Open(config.DatabaseConfig.DatabaseReadURL)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("DATABASE CONNECTED SUCCESSFULLY")

	rdb := cacheServices.InitRedisCacheService()
	cacheService := cacheServices.NewRedisCacheService(rdb)

	res, err := cacheService.CheckRedisConnection()

	if err != nil {
		fmt.Println("Redis not connected properly", err)
		return
	} else {
		fmt.Println("Redis connected succesfully....", res)
	}

	bbDelivery.NewBBHandler(e, bbUsecase.NewUser(bbRepository.NewUser(db)))
	log.Fatal(e.Start(":" + "80"))

}
