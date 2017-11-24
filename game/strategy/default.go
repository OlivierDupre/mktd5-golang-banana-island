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

	printMap(helper)

	//gameState := helper.GameState()
	//gameState.Map.Cell()
	myPosition := WhereIAm(helper)
	fmt.Println("I am here: ", myPosition)

	aroundMe := WhatsAroundMe(myPosition.x, myPosition.y, helper.GameState().Map)
	fmt.Print("North: ", aroundMe[mediator.North])
	fmt.Print("East: ", aroundMe[mediator.East])
	fmt.Print("South: ", aroundMe[mediator.South])
	fmt.Print("West: ", aroundMe[mediator.West])

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

type Position struct {
	x   int
	y   int
}
func WhereIAm(helper Helper) Position{
	var maps [][]mediator.Cell = helper.GameState().Map
	for irow := range maps {
		for icol := range maps[irow] {
			if helper.IsMe(maps[irow][icol]) {
				return  Position{irow, icol}
			}
		}
	}
	panic("I'm lost !")
}

func WhatsAroundMe(x int, y int, maps mediator.Map) map[mediator.Direction]mediator.Cell {
	aroundme := make(map[mediator.Direction]mediator.Cell)

	var err error

	aroundme[mediator.North], err = maps.Cell(x, y-1)
	if err != nil {
		aroundme[mediator.North] = 2
	}

	aroundme[mediator.South], err = maps.Cell(x, y+1)
	if err != nil {
		aroundme[mediator.South] = 2
	}

	aroundme[mediator.West], err = maps.Cell(x-1,y)
	if err != nil {
		aroundme[mediator.West] = 2
	}

	aroundme[mediator.East], err = maps.Cell(x+1,y)
	if err != nil {
		aroundme[mediator.East] = 2
	}

	return aroundme
}