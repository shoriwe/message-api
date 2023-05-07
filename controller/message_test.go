package controller

import (
	"testing"

	"github.com/shoriwe/message-api/common/random"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"github.com/stretchr/testify/assert"
)

func testSendMessage(t *testing.T, c *Controller) {
	t.Run("Valid", func(tt *testing.T) {
		// Initialize FCM

		//
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		u2 := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u2).Error)
		d2 := models.RandomDevice(u2)
		assert.Nil(tt, c.DB.Create(d2).Error)
		msg := models.Message{
			SenderUUID:    u.UUID,
			RecipientUUID: u2.UUID,
			Title:         random.String(),
			Body:          random.String(),
		}
		assert.Nil(tt, c.SendMessage(s, &msg))
	})
}

func TestSendMessage(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		c := NewMem()
		defer c.Close()
		testSendMessage(tt, c)
	})
}
