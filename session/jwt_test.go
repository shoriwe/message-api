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
		ss, tok, err := j.Parse(jwtTok)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		assert.Equal(tt, s.UserUUID, ss.UserUUID)
		assert.Equal(tt, s.FirebaseToken, ss.FirebaseToken)
	})
}

func TestJWT_Parse(t *testing.T) {
	t.Run("Valid", func(tt *testing.T) {
		j := NewDefault()
		s := NewSession(uuid.NewV4(), random.String())
		jwtTok := j.New(s)
		assert.Greater(tt, len(jwtTok), 0)
		ss, tok, err := j.Parse(jwtTok)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		assert.Equal(tt, s.UserUUID, ss.UserUUID)
		assert.Equal(tt, s.FirebaseToken, ss.FirebaseToken)
	})
	t.Run("Invalid", func(tt *testing.T) {
		j := NewDefault()
		_, _, err := j.Parse("{}}")
		assert.NotNil(tt, err)
	})
	t.Run("No firebaseToken", func(tt *testing.T) {
		j := NewDefault()
		jwtTok := j.New(jwt.MapClaims{"userUUID": uuid.NewV4().String()})
		assert.Greater(tt, len(jwtTok), 0)
		_, _, err := j.Parse(jwtTok)
		assert.NotNil(tt, err)
	})
	t.Run("Invalid firebaseToken", func(tt *testing.T) {
		j := NewDefault()
		jwtTok := j.New(jwt.MapClaims{"firebaseToken": 1, "userUUID": uuid.NewV4().String()})
		assert.Greater(tt, len(jwtTok), 0)
		_, _, err := j.Parse(jwtTok)
		assert.NotNil(tt, err)
	})
	t.Run("No userUUID", func(tt *testing.T) {
		j := NewDefault()
		jwtTok := j.New(jwt.MapClaims{"firebaseToken": random.String()})
		assert.Greater(tt, len(jwtTok), 0)
		_, _, err := j.Parse(jwtTok)
		assert.NotNil(tt, err)
	})
	t.Run("Invalid userUUID", func(tt *testing.T) {
		j := NewDefault()
		jwtTok := j.New(jwt.MapClaims{"firebaseToken": random.String(), "userUUID": 1.0})
		assert.Greater(tt, len(jwtTok), 0)
		_, _, err := j.Parse(jwtTok)
		assert.NotNil(tt, err)
	})
}
