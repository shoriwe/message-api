package session

import (
	"testing"

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
