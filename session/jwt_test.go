package session

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/message-api/common/random"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	t.Run("Create", func(tt *testing.T) {
		j := NewDefault()
		s := NewSession(uuid.NewV4(), random.String())
		jwtTok := j.New(s)
		assert.Greater(tt, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		assert.Equal(tt, s.UserUUID.String(), tok.Claims.(jwt.MapClaims)["userUUID"])
		assert.Equal(tt, s.FirebaseToken, tok.Claims.(jwt.MapClaims)["firebaseToken"])
	})
}
