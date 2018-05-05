package tests

import (
	"github.com/MetalRex101/auth-server/app/db"
	"github.com/MetalRex101/auth-server/config"
	"log"
	"github.com/jinzhu/gorm"
)

func withTx(callback func(dbConn *gorm.DB) error) {
	conf := config.GetConfig("testing")

	tx := db.Init(conf).Begin()
	if err := tx.Error; err != nil {
		log.Panic(err)
	}

	callbackError := callback(tx)

	if err := tx.Rollback().Error; err != nil {
		log.Panic(err)
	}

	if callbackError != nil {
		log.Panic(callbackError)
	}
}
