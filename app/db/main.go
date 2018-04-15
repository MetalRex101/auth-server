package db

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/config"
	"log"
)

var Gorm *gorm.DB

func init() {
	var err error
	conf := config.Instance
	uri := conf.DB.DSN

	Gorm, err = gorm.Open(conf.DB.Server, uri)

	if err != nil {
		log.Panicf("Could not connect to database: %s", err)
	}
}
