package db

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/config"
	"log"
	"github.com/rubenv/sql-migrate"
	"fmt"
)

func Init(config *config.Config) *gorm.DB {
	uri := config.DB.DSN

	conn, err := gorm.Open(config.DB.Server, uri)

	if err != nil {
		log.Panicf("Could not connect to database: %s", err)
	}

	return conn
}

func MigrateUp(db *gorm.DB, config *config.Config) {
	migrations := &migrate.FileMigrationSource{
		Dir: config.App.MigrationDir,
	}

	n, err := migrate.Exec(db.DB(), config.DB.Server, migrations, migrate.Up)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Migration failed: %s", err))
	}

	fmt.Printf("Applied %d migrations!\n", n)
}

func MigrateDown (db *gorm.DB, config *config.Config) {
	migrations := &migrate.FileMigrationSource{
		Dir: config.App.MigrationDir,
	}

	n, err := migrate.Exec(db.DB(), config.DB.Server, migrations, migrate.Down)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Migration Rollback failed: %s", err))
	}

	fmt.Printf("Rollbacked %d migrations!\n", n)
}