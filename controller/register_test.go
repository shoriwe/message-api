package controller

import (
	"testing"

	"github.com/shoriwe/message-api/models"
	"github.com/stretchr/testify/assert"
)

func testRegister(t *testing.T, c *Controller) {
	t.Run("Valid Registration", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.Register(u))
		var user models.User
		assert.Nil(tt, c.DB.Where("email = ?", u.Email).First(&user).Error)
		assert.Equal(tt, u.Email, user.Email)
	})
}

func TestRegister(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		c := NewMem()
		defer c.Close()
		testRegister(tt, c)
	})
}
