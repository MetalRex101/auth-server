package db

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/config"
	"log"
)

func Init() *gorm.DB {
	conf := config.Instance
	uri := conf.DB.DSN

	conn, err := gorm.Open(conf.DB.Server, uri)

	if err != nil {
		log.Panicf("Could not connect to database: %s", err)
	}

	return conn
}