package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomDevice(t *testing.T) {
	assert.NotEqual(t, RandomDevice(RandomUser()).FirebaseToken, RandomDevice(RandomUser()).FirebaseToken)
}
