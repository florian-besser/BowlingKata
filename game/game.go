package game

type Game struct {
	rolls      [21]int
	rollsIndex int
	sum        int
}

func (game *Game) GetScore() int {
	sum := 0
	for i := 0; i < 21; i++ {
		sum += game.rolls[i]
	}
	return sum
}

func (game *Game) Roll(roll int) {
	game.rolls[game.rollsIndex] = roll
	game.rollsIndex++
	game.sum += roll
}
