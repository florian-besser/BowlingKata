package game

type Game struct {
	rolls      [21]int
	rollsIndex int
}

func (game *Game) GetScore() int {
	sum := 0
	index := 0
	for frame := 0; frame < 10; frame++ {
		isSpare := game.rolls[index]+game.rolls[index+1] == 10
		if isSpare {
			sum += game.rolls[index] + game.rolls[index+1] + game.rolls[index+2]
		} else {
			sum += game.rolls[index] + game.rolls[index+1]
		}
		index += 2
	}
	return sum
}

func (game *Game) Roll(roll int) {
	game.rolls[game.rollsIndex] = roll
	game.rollsIndex++
}
