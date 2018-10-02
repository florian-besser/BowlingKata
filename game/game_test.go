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

func Test_Strike_NextTwoRollsCountTwice(t *testing.T) {
	g := Game{}

	g.Roll(10)
	g.Roll(4)
	g.Roll(3)

	assert.Equal(t, g.GetScore(), 24)
}

// Whoops! It looks like the last frame was not a special case at all, this test is green anyway
func Test_LastFrame_AllowThreeRolls(t *testing.T) {
	g := Game{rollsIndex: 18}

	g.Roll(10)
	g.Roll(4)
	g.Roll(3)

	assert.Equal(t, g.GetScore(), 17)
}

func TestEvil_StrikeButNotSpare_FrameRulesApply(t *testing.T) {
	g := &Game{}

	g.Roll(10)
	g.Roll(0)
	g.Roll(1)

	// Some people may implement computation of strike and spare as an "or", meaning the 10 will count as a strike (correct) and also as a "spare" (wrong)
	assert.Equal(t, g.GetScore(), 12)
}

func TestEvil_GameAsPerSlideDeck_CorrectResult(t *testing.T) {
	g := &Game{}

	// This test acts as a "God test" meaning it tests many different things together. This makes it hard to debug and brittle if you break something.
	// Avoid such large tests as this one and prefer smaller, more precise tests as these are more robust and easier to debug.
	g.Roll(1)
	g.Roll(4)
	g.Roll(4)
	g.Roll(5)
	g.Roll(6)
	g.Roll(4)
	g.Roll(5)
	g.Roll(5)
	g.Roll(10)
	g.Roll(0)
	g.Roll(1)
	g.Roll(7)
	g.Roll(3)
	g.Roll(6)
	g.Roll(4)
	g.Roll(10)
	g.Roll(2)
	g.Roll(8)
	g.Roll(6)

	assert.Equal(t, g.GetScore(), 133)
}

func TestEvil_Perfect(t *testing.T) {
	g := &Game{}

	// This test also tests many things together, but it could also uncover a bug in the computation of strike bonuses
	// Some participants implement frames and will have a boolean flag to carry over if the previous frame contained a
	// spare or strike, in order to count some rolls of the next frame twice. But this could be detected with a simpler
	// test that simply rolls three strikes in a row!
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)
	g.Roll(10)

	assert.Equal(t, g.GetScore(), 300)
}
