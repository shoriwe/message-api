package controller

import (
	"testing"

	"github.com/shoriwe/message-api/models"
	"github.com/stretchr/testify/assert"
)

func testDownloadPicture(t *testing.T, c *Controller) {
	t.Run("Valid Registration", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		contents, dErr := c.DownloadPicture(u.UUID.String())
		assert.Nil(tt, dErr)
		assert.Equal(tt, models.TestingImage, string(contents))
	})
}

func TestDownloadPicture(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		c := NewMem()
		defer c.Close()
		testDownloadPicture(tt, c)
	})
}
