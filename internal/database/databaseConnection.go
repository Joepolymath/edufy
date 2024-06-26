package database

import (
	"Learnium/internal/config"
	"strings"

	// "Learnium/internal/pkg/logger"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cfg = config.Config // get config model as local variable within database module

func DBConnection() *gorm.DB {
	/*
		This is used to connect to the database, and if the database has issues, i panic on the console
	*/

	var connectionString string

	if strings.ToLower(cfg.ENVIRON) != "development" {
		connectionString = fmt.Sprintf(" host=%s user=%s password=%s dbname=%s port=%d sslmode=require", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabaseName, cfg.PostgresPort)
	} else {
		connectionString = fmt.Sprintf(" host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabaseName, cfg.PostgresPort)
	}

	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panicln("Error connecting to database:::", err)
	}
	log.Println("Connected successfully to database as user: ", cfg.PostgresUser)

	return connection
}

func CloseDB(db *gorm.DB) {

	closeDb, err := db.DB()
	if err != nil {
		log.Panicln("Error closing database connection::: ", err)
	}

	err = closeDb.Close()
	if err != nil {
		log.Panicln(err)
	}
}
