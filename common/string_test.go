package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	var length int
	length = 32
	assert.Equal(t, len(RandomString(length)), length)
	length = 89
	assert.Equal(t, len(RandomString(length)), length)
	length = 1
	assert.Equal(t, len(RandomString(length)), length)
	length = 9999
	assert.Equal(t, len(RandomString(length)), length)

	assert.Equal(t, RandomString(0), "")
	assert.Equal(t, RandomString(-1), "")
}
