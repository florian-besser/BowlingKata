package game

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_NewGame_ScoreIs0(t *testing.T) {
	g := Game{}

	assert.Equal(t, g.GetScore(), 0)
}

func Test_TwoRolls_ScoreIsSum(t *testing.T) {
	g := Game{}

	g.Roll(3)
	g.Roll(4)

	assert.Equal(t, g.GetScore(), 7)
}

func Test_Spare_NextRollCountsTwice(t *testing.T) {
	g := Game{}

	g.Roll(3)
	g.Roll(7)
	g.Roll(6)

	assert.Equal(t, g.GetScore(), 22)
}
