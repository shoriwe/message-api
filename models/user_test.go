package models

import (
	"testing"

	"github.com/shoriwe/message-api/common/sqlite"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRandomUser(t *testing.T) {
	assert.NotEqual(t, RandomUser(), RandomUser())
}

func testUser(t *testing.T, db *gorm.DB) {
	t.Run("ValidUser", func(tt *testing.T) {
		u := RandomUser()
		assert.Nil(tt, db.Create(u).Error)
	})
	t.Run("RepeatedUser", func(tt *testing.T) {
		u := RandomUser()
		assert.Nil(tt, db.Create(u).Error)
		assert.NotNil(tt, db.Create(u).Error)
	})
	t.Run("Claims", func(tt *testing.T) {
		u := RandomUser()
		assert.Nil(tt, db.Create(u).Error)
		assert.NotNil(tt, u.Claims())
	})
	t.Run("FromClaims-Valid", func(tt *testing.T) {
		u := RandomUser()
		var u2 User
		assert.Nil(tt, u2.FromClaims(u.Claims()))
		assert.Equal(tt, u.UUID, u2.UUID)
	})
	t.Run("FromClaims-NoUUID", func(tt *testing.T) {
		u := RandomUser()
		var u2 User
		claims := u.Claims()
		delete(claims, "uuid")
		assert.NotNil(tt, u2.FromClaims(claims))
	})
	t.Run("Authenticate-Succeed", func(tt *testing.T) {
		u := RandomUser()
		u.PasswordHash, _ = bcrypt.GenerateFromPassword([]byte(u.Password), DefaultPasswordCost)
		assert.True(tt, u.Authenticate(u.Password))
	})
	t.Run("Authenticate-Fail", func(tt *testing.T) {
		u := RandomUser()
		u.PasswordHash, _ = bcrypt.GenerateFromPassword([]byte(u.Password), DefaultPasswordCost)
		assert.False(tt, u.Authenticate("wrong"))
	})
	t.Run("BeforeSave-ValidUser", func(tt *testing.T) {
		u := RandomUser()
		assert.Nil(tt, db.Create(u).Error)
	})
	t.Run("BeforeSave-NoPassword", func(tt *testing.T) {
		u := RandomUser()
		u.Password = ""
		assert.NotNil(tt, db.Create(u).Error)
	})
	t.Run("BeforeSave-InvalidPhone", func(tt *testing.T) {
		u := RandomUser()
		u.PhoneNumber = ""
		assert.NotNil(tt, db.Create(u).Error)
	})
	t.Run("BeforeSave-InvalidEmail", func(tt *testing.T) {
		u := RandomUser()
		u.Email = ""
		assert.NotNil(tt, db.Create(u).Error)
	})
}

func TestUser(t *testing.T) {
	t.Run("SQLite", func(tt *testing.T) {
		db := sqlite.NewMem()
		conn, _ := db.DB()
		defer conn.Close()
		assert.Nil(tt, db.AutoMigrate(&User{}, &Device{}, &Message{}))
		testUser(tt, db)
	})
}
