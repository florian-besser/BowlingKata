package game

type Game struct {
	sum int
}

func (game *Game) GetScore() int {
	return game.sum
}

func (game *Game) Roll(roll int) {
	game.sum += roll
}
