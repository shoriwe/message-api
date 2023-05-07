package controller

import (
	"github.com/shoriwe/message-api/common/sqlite"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"gorm.io/gorm"
)

type Controller struct {
	DB  *gorm.DB
	JWT *session.JWT
}

func (c *Controller) Close() error {
	conn, err := c.DB.DB()
	if err != nil {
		return err
	}
	return conn.Close()
}

func New(db *gorm.DB, j *session.JWT) *Controller {
	db.AutoMigrate(
		&models.User{}, &models.Device{}, &models.Message{},
	)
	return &Controller{
		DB:  db,
		JWT: j,
	}
}

func NewMem() *Controller {
	return New(sqlite.NewMem(), session.NewDefault())
}
