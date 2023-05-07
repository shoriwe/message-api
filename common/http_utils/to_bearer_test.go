package http_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBearer(t *testing.T) {
	assert.Equal(t, "Bearer WUVZ", ToBearer("YEY"))
}
