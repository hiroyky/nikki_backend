package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeID(t *testing.T) {
	actual := EncodeGraphQLID("name", 1)
	expected := "bmFtZS0x"
	assert.Equal(t, expected, actual)
}

func TestDecodeID(t *testing.T) {
	actual, err := DecodeGraphQLID("bmFtZS0x")
	expected := 1
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
