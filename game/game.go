package game

type Game struct {
	rolls      [21]int
	rollsIndex int
}

func (game *Game) GetScore() int {
	sum := 0
	// TODO: Index only until 19 is wrong!
	for i := 0; i < 19; i++ {
		// TODO this is not the correct way to detect a spare!
		isSpare := game.rolls[i]+game.rolls[i+1] == 10
		if isSpare {
			// This is not a good way to add the bonus, hard to understand!
			sum += game.rolls[i+2]
		}
		sum += game.rolls[i]
	}
	return sum
}

func (game *Game) Roll(roll int) {
	game.rolls[game.rollsIndex] = roll
	game.rollsIndex++
}
