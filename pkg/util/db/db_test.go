package dbutil

import (
	"sync"
	"testing"

	"github.com/Storm-cpu/core/config"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func mockConfig() config.Configuration {
	return config.Configuration{
		DbPsn: "postgres://localhost:5432/postgres?user=postgres&password=core123",
	}
}

func TestConnectDB(t *testing.T) {
	cfg := mockConfig()

	db := ConnectDB(cfg)

	assert.NotNil(t, db, "Database connection should not be nil")

	pgDB, err := db.DB()
	assert.NoError(t, err, "Should not error when getting postgres database")
	assert.NoError(t, pgDB.Ping(), "Database should be reachable")
}

func TestConnectDBSingleton(t *testing.T) {
	cfg := mockConfig()
	var wg sync.WaitGroup

	var db1, db2 *gorm.DB

	wg.Add(1)

	go func() {
		defer wg.Done()
		db1 = ConnectDB(cfg)
	}()

	db2 = ConnectDB(cfg)

	wg.Wait()

	assert.NotNil(t, db1, "First database connection should not be nil")
	assert.NotNil(t, db2, "Second database connection should not be nil")
	assert.Equal(t, db1, db2, "Both connections should point to the same instance")
}
