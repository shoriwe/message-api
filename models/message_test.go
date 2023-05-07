package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomMessage(t *testing.T) {
	assert.NotEqual(t,
		RandomMessage(RandomUser(), RandomUser()).Body,
		RandomMessage(RandomUser(), RandomUser()).Body,
	)
}
