package controller

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/shoriwe/message-api/common/random"
	"github.com/shoriwe/message-api/common/sqlite"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type Controller struct {
	DB     *gorm.DB
	JWT    *session.JWT
	App    *firebase.App
	Client *messaging.Client
}

func (c *Controller) Close() error {
	conn, err := c.DB.DB()
	if err != nil {
		return err
	}
	return conn.Close()
}

func New(db *gorm.DB, j *session.JWT, app *firebase.App) *Controller {
	db.AutoMigrate(
		&models.User{}, &models.Device{}, &models.Message{},
		&models.MessageResponse{},
	)
	client, cErr := app.Messaging(context.Background())
	if cErr != nil {
		panic(cErr)
	}
	return &Controller{
		DB:     db,
		JWT:    j,
		App:    app,
		Client: client,
	}
}

func NewMem() *Controller {
	config := firebase.Config{
		ProjectID: random.String(),
	}
	app, nErr := firebase.NewApp(context.Background(), &config, option.WithoutAuthentication())
	if nErr != nil {
		panic(nErr)
	}
	return New(sqlite.NewMem(), session.NewDefault(), app)
}
