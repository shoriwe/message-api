package controller

import (
	"testing"

	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"github.com/stretchr/testify/assert"
)

func testListUsers(t *testing.T, c *Controller) {
	t.Run("Simple List", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		// Init users
		for i := 0; i < 10; i++ {
			u2 := models.RandomUser()
			assert.Nil(tt, c.Register(u2))
		}
		filter := Filter[models.User]{
			PageSize: 10,
		}
		users, fErr := c.ListUsers(s, &filter)
		assert.Nil(tt, fErr)
		assert.Len(tt, users, 10)
	})
	t.Run("By Email", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		// Init user
		u2 := models.RandomUser()
		assert.Nil(tt, c.Register(u2))
		filter := Filter[models.User]{
			PageSize: 10,
			Target: models.User{
				Email: u2.Email,
			},
		}
		users, fErr := c.ListUsers(s, &filter)
		assert.Nil(tt, fErr)
		assert.Len(tt, users, 1)
	})
	t.Run("By Name", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		// Init user
		u2 := models.RandomUser()
		assert.Nil(tt, c.Register(u2))
		filter := Filter[models.User]{
			PageSize: 10,
			Target: models.User{
				Name: u2.Name,
			},
		}
		users, fErr := c.ListUsers(s, &filter)
		assert.Nil(tt, fErr)
		assert.Len(tt, users, 1)
	})
	t.Run("By Job", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		// Init user
		u2 := models.RandomUser()
		assert.Nil(tt, c.Register(u2))
		filter := Filter[models.User]{
			PageSize: 10,
			Target: models.User{
				Job: u2.Job,
			},
		}
		users, fErr := c.ListUsers(s, &filter)
		assert.Nil(tt, fErr)
		assert.Len(tt, users, 1)
	})
	t.Run("By PhoneNumber", func(tt *testing.T) {
		u := models.RandomUser()
		assert.Nil(tt, c.DB.Create(u).Error)
		d := models.RandomDevice(u)
		assert.Nil(tt, c.DB.Create(d).Error)
		s := session.NewSession(u.UUID, d.FirebaseToken)
		// Init user
		u2 := models.RandomUser()
		assert.Nil(tt, c.Register(u2))
		filter := Filter[models.User]{
			Target: models.User{
				PhoneNumber: u2.PhoneNumber,
			},
		}
		users, fErr := c.ListUsers(s, &filter)
		assert.Nil(tt, fErr)
		assert.Len(tt, users, 1)
	})
}

func TestListUsers(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		c := NewMem()
		defer c.Close()
		testListUsers(tt, c)
	})
}
