package models

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/message-api/common/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

/*
TestBaseModel basic unit test to reduce the coverage footprint
*/
func testBaseModel_BeforeSafe(t *testing.T, db *gorm.DB) {
	db.AutoMigrate(&Model{})
	t.Run("Null UUID", func(t *testing.T) {
		assert.Nil(t, db.Create(&Model{}).Error)
	})
	t.Run("Set UUID", func(t *testing.T) {
		assert.Nil(t, db.Create(&Model{UUID: uuid.NewV4()}).Error)
	})
}

func TestBaseModel_BeforeSafe(t *testing.T) {
	t.Run("SQLite", func(t *testing.T) {
		db := sqlite.NewMem()
		conn, _ := db.DB()
		defer conn.Close()
		testBaseModel_BeforeSafe(t, db)
	})
}
