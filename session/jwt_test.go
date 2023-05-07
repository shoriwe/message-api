package session

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/message-api/models"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	t.Run("New", func(tt *testing.T) {
		j := New([]byte("TESTING"))
		jwtTok := j.New(jwt.MapClaims{"uuid": "my_uuid"})
		assert.Greater(tt, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
	})
	t.Run("User", func(tt *testing.T) {
		j := New([]byte("TESTING"))
		user := models.RandomUser()
		user.UUID = uuid.NewV4()
		jwtTok := j.New(user.Claims())
		assert.Greater(tt, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		var user2 models.User
		assert.Nil(tt, user2.FromClaims(tok.Claims.(jwt.MapClaims)))
		assert.Equal(tt, user.UUID, user2.UUID)
	})
}
