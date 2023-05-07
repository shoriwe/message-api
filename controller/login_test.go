package controller

import (
	"testing"

	"github.com/shoriwe/message-api/common/random"
	"github.com/shoriwe/message-api/models"
	"github.com/stretchr/testify/assert"
)

func testLogin(t *testing.T, c *Controller) {
	t.Run("Valid Credentials", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		token := random.String()
		user, _, loginErr := c.Login(&Login{Email: *u.Email, Password: *u.Password, FirebaseToken: token})
		assert.Nil(tt, loginErr)
		assert.Equal(tt, u.UUID, user.UUID)
		var device models.Device
		assert.Nil(tt, c.DB.Where("firebase_token = ?", token).First(&device).Error)
		assert.Equal(tt, token, device.FirebaseToken)
	})
	t.Run("Invalid Email", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		token := random.String()
		user, _, loginErr := c.Login(&Login{Email: "INVALID", Password: *u.Password, FirebaseToken: token})
		assert.NotNil(tt, loginErr)
		assert.Nil(tt, user)
	})
	t.Run("Invalid Password", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		token := random.String()
		user, _, loginErr := c.Login(&Login{Email: *u.Email, Password: "INVALID", FirebaseToken: token})
		assert.NotNil(tt, loginErr)
		assert.Nil(tt, user)
	})
}

func TestLogin(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		c := NewMem()
		defer c.Close()
		testLogin(tt, c)
	})
}
