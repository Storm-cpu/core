package dbutil

import (
	"core/config"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	instance bool
	lock     sync.Mutex
)

func ConnectDB(cfg config.Configuration) *gorm.DB {
	lock.Lock()
	defer lock.Unlock()

	if instance {
		fmt.Println("Database already connected!")
		return db
	}

	var err error
	db, err = gorm.Open(postgres.Open(cfg.DbPsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Connected to Database!")
	instance = true

	return db
}
