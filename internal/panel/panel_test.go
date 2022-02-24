package panel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	got := getTrue()
	assert.Equal(t, got, true)
}
