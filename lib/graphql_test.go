package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeID(t *testing.T) {
	actual := EncodeGraphQLID("name", 1)
	expected := "bmFtZTox"
	assert.Equal(t, expected, actual)
}

func TestDecodeID(t *testing.T) {
	actual, err := DecodeGraphQLID("bmFtZTox")
	expected := 1
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
