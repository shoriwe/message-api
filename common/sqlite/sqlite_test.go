package sqlite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLite(t *testing.T) {
	t.Run("ValidDSN", func(tt *testing.T) {
		db := New(":memory:")
		conn, err := db.DB()
		assert.Nil(tt, err)
		assert.Nil(tt, conn.Close())
	})
}

func TestSQLite_NewMem(t *testing.T) {
	t.Run("ValidDSN", func(tt *testing.T) {
		db := NewMem()
		conn, err := db.DB()
		assert.Nil(tt, err)
		assert.Nil(tt, conn.Close())
	})
}
