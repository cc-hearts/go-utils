package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidJson(t *testing.T) {
	assert.True(t, IsValidJson(`{"name": "John", "age": 30}`))
	assert.False(t, IsValidJson(`{"name": "John", "age": 30,, "city": "New York"}`))
}
