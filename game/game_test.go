package game

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFirst(t *testing.T) {
	g := Game{}

	assert.Equal(t, g.GetScore(), 0)
}
