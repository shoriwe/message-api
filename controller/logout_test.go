package controller

import (
	"testing"

	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"github.com/stretchr/testify/assert"
)

func testLogout(t *testing.T, c *Controller) {
	t.Run("Valid Token", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		assert.Nil(tt, c.Logout(s))
	})
	t.Run("Firebase token of other user", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		u2 := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u2).Error)
		d := models.RandomDevice(u2)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		assert.NotNil(tt, c.Logout(s))
	})
}

func TestLogout(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		c := NewMem()
		defer c.Close()
		testLogout(tt, c)
	})
}
