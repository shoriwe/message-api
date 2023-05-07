package session

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/message-api/common/random"
	"github.com/stretchr/testify/assert"
)

func TestSession(t *testing.T) {
	t.Run("GetExpirationTime", func(tt *testing.T) {
		s := NewSession(uuid.NewV4(), random.String())
		_, err := s.GetExpirationTime()
		assert.Nil(tt, err)
	})
	t.Run("GetIssuedAt", func(tt *testing.T) {
		s := NewSession(uuid.NewV4(), random.String())
		_, err := s.GetIssuedAt()
		assert.Nil(tt, err)
	})
	t.Run("GetNotBefore", func(tt *testing.T) {
		s := NewSession(uuid.NewV4(), random.String())
		_, err := s.GetNotBefore()
		assert.Nil(tt, err)
	})
	t.Run("GetIssuer", func(tt *testing.T) {
		s := NewSession(uuid.NewV4(), random.String())
		_, err := s.GetIssuer()
		assert.Nil(tt, err)
	})
	t.Run("GetSubject", func(tt *testing.T) {
		s := NewSession(uuid.NewV4(), random.String())
		_, err := s.GetSubject()
		assert.Nil(tt, err)
	})
	t.Run("GetAudience", func(tt *testing.T) {
		s := NewSession(uuid.NewV4(), random.String())
		_, err := s.GetAudience()
		assert.Nil(tt, err)
	})
}

func TestNewSessionFromToken(t *testing.T) {
	t.Run("Valid", func(tt *testing.T) {
		j := NewDefault()
		s := NewSession(uuid.NewV4(), random.String())
		jwtTok := j.New(s)
		assert.Greater(tt, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		ss, sErr := NewSessionFromToken(tok)
		assert.Nil(tt, sErr)
		assert.Equal(tt, s.UserUUID, ss.UserUUID)
		assert.Equal(tt, s.FirebaseToken, ss.FirebaseToken)
	})
	t.Run("No firebaseToken", func(tt *testing.T) {
		j := NewDefault()
		jwtTok := j.New(jwt.MapClaims{"userUUID": uuid.NewV4().String()})
		assert.Greater(tt, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		_, sErr := NewSessionFromToken(tok)
		assert.NotNil(tt, sErr)
	})
	t.Run("Invalid firebaseToken", func(tt *testing.T) {
		j := NewDefault()
		jwtTok := j.New(jwt.MapClaims{"firebaseToken": 1, "userUUID": uuid.NewV4().String()})
		assert.Greater(tt, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		_, sErr := NewSessionFromToken(tok)
		assert.NotNil(tt, sErr)
	})
	t.Run("No userUUID", func(tt *testing.T) {
		j := NewDefault()
		jwtTok := j.New(jwt.MapClaims{"firebaseToken": random.String()})
		assert.Greater(tt, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		_, sErr := NewSessionFromToken(tok)
		assert.NotNil(tt, sErr)
	})
	t.Run("Invalid userUUID", func(tt *testing.T) {
		j := NewDefault()
		jwtTok := j.New(jwt.MapClaims{"firebaseToken": random.String(), "userUUID": 1.0})
		assert.Greater(tt, len(jwtTok), 0)
		tok, err := jwt.Parse(jwtTok, j.KeyFunc)
		assert.Nil(tt, err)
		assert.True(tt, tok.Valid)
		_, sErr := NewSessionFromToken(tok)
		assert.NotNil(tt, sErr)
	})
}
