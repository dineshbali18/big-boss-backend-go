package main

import (
	"big-boss-7/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func init() {
	//Initialize config
	config.InitializeConfig()
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
}
