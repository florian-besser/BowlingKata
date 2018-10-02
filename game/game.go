package game

type Game struct {
	rolls      [21]int
	rollsIndex int
}

func (game *Game) GetScore() int {
	sum := 0
	index := 0
	for frame := 0; frame < 10; frame++ {
		sum += game.rolls[index]
		index++
	}
	return sum
}

func (game *Game) Roll(roll int) {
	game.rolls[game.rollsIndex] = roll
	game.rollsIndex++
}
