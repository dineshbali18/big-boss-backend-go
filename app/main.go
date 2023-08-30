package main

import (
	"big-boss-7/config"
	"fmt"
	"log"

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
		log.Fatal(err.Error())
	}

	// Establish data base connection
	db, err := gorm.Open(mysql.Open(config.DatabaseConfig.DatabaseURL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Specifying DB Reader and Writer
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(config.DatabaseConfig.DatabaseWriteURL)},
		Replicas: []gorm.Dialector{mysql.Open(config.DatabaseConfig.DatabaseReadURL)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("DATABASE CONNECTED SUCCESSFULLY")

	bbDelivery.NewBBHandler(e, bbUsecase.NewUser(bbRepository.NewUser(db)))
	log.Fatal(e.Start(":" + "3306"))

}
