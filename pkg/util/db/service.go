package dbutil

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dbPsn string, cfg *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbPsn), cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}
