/*
sqlite will be only used for unit tests and when no postgres DSN is provided
*/
package sqlite

import (
	"fmt"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var currentMemDB = 0

func NewMem() *gorm.DB {
	currentMemDB++
	return New(fmt.Sprintf("file:test-%d?mode=memory&cache=shared", currentMemDB))
}

func New(file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	return db
}
