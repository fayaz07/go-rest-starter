package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContainsValid(t *testing.T) {
	assert := assert.New(t)

	var mockSlice = []string{"a", "b", "c", "apple", "banana", "aeroplane"}

	assert.True(Contains(mockSlice, "a"))
	assert.True(Contains(mockSlice, "apple"))
	assert.True(Contains(mockSlice, "aeroplane"))

	assert.False(Contains(mockSlice, "ab"))
	assert.False(Contains(mockSlice, "banan"))
}
