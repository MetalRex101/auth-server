package db

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/config"
	"log"
)

func Init(config *config.Config) *gorm.DB {
	uri := config.DB.DSN

	conn, err := gorm.Open(config.DB.Server, uri)

	if err != nil {
		log.Panicf("Could not connect to database: %s", err)
	}

	return conn
}