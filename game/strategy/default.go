package strategy

import (
	"mktd5/mktd-island/client/log"
	"mktd5/mktd-island/client/mediator"
	"math/rand"
	"fmt"
)

type DefaultMoveStrategy struct {
	Logger log.LoggerInterface `inject:""`
}

// DecideWhereToGo is invoked when it is the player's turn to make a move.
// Given a helper object that can provide useful context, it has to decide (in a limited
// time frame) what will be the next move for the player.
// The function must return one of the following directions:
//   - mediator.North
//   - mediator.East
//   - mediator.South
//   - mediator.West
//   - mediator.None
// It can also return an error. If an error is returned, the player will make no move
// for the current turn and the error will be logged.
func (d *DefaultMoveStrategy) DecideWhereToGo(helper Helper) (mediator.Direction, error) {
	d.Logger.Info("let's go!", nil)

	//rand.Seed(42) // If you fix this number, the random will always start at the same number
	newDirection := randomMove()

	//gameState := helper.GameState()
	printMap(helper)

	return newDirection, nil
}

func randomMove() mediator.Direction{
	directions := []mediator.Direction{
		mediator.North,
		mediator.East,
		mediator.South,
		mediator.West,
	}

	return directions[rand.Intn(len(directions))]
}

func printMap(helper Helper) {
	state := helper.GameState()
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			cell, _ := state.Map.Cell(x, y)
			if cell.Banana() {
				fmt.Print("b")
			} else if helper.IsMe(cell) {
				fmt.Print("M")
			} else if cell.Player() {
				fmt.Print("m")
			} else if cell.Wall() {
				fmt.Print(".")
			} else if cell.Empty() {
				fmt.Print(" ")
			} else {
				fmt.Print("?")
			}
		}
		fmt.Println("")
	}
}