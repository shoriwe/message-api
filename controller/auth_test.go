package controller

import (
	"testing"

	"github.com/shoriwe/message-api/common/random"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"github.com/stretchr/testify/assert"
)

func testAuthenticate(t *testing.T, c *Controller) {
	t.Run("Authorized", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		user, aErr := c.Authenticate(s)
		assert.Nil(tt, aErr)
		assert.Equal(tt, u.UUID, user.UUID)
	})
	t.Run("Unauthorized", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, random.String())
		_, aErr := c.Authenticate(s)
		assert.NotNil(tt, aErr)
	})
}

func TestAuthenticate(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		c := NewMem()
		defer c.Close()
		testAuthenticate(tt, c)
	})
}
