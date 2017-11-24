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

	var actualPos Position = whereIAm(helper)
	fmt.Printf("Actual position %v", actualPos)

	fmt.Println("Current map")
	printMap(helper)

	var aroundMe map[mediator.Direction]mediator.Cell = WhatsAroundMe(actualPos,helper.GameState().Map)
	fmt.Print("North: ", aroundMe[mediator.North])
	fmt.Print("East: ", aroundMe[mediator.East])
	fmt.Print("South: ", aroundMe[mediator.South])
	fmt.Print("West: ", aroundMe[mediator.West])

	dir := chooseDirection(aroundMe)

	fmt.Print(actualPos.X, actualPos.Y)
	return dir, nil
}

type Position struct {
	X  int
	Y  	int
}

func chooseDirection(aroundMe map[mediator.Direction]mediator.Cell) mediator.Direction {
	dir := mediator.North
	if aroundMe[mediator.North].Banana() {
		dir = mediator.North
	} else if aroundMe[mediator.East].Banana() {
		dir = mediator.East
	} else if aroundMe[mediator.West].Banana() {
		dir = mediator.West
	} else if aroundMe[mediator.South].Banana() {
		dir = mediator.South
	} else {
		dir = randomMove()
	}

	return dir
}

func whereIAm(helper Helper) Position{
	var maps [][]mediator.Cell = helper.GameState().Map
	for irow := range maps {
		for icol := range maps[irow] {
			if helper.IsMe(maps[irow][icol]) {
				return  Position{icol, irow}
			}
		}
	}
	panic("I'm lost !")
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

func WhatsAroundMe(pos Position, state mediator.Map) map[mediator.Direction]mediator.Cell {

	aroundme := make(map[mediator.Direction]mediator.Cell)
	var err error
	aroundme[mediator.North], err = state.Cell(pos.X, pos.Y-1)
	if err != nil {
		aroundme[mediator.North] = 2
	}

	aroundme[mediator.South], err = state.Cell(pos.X, pos.Y+1)
	if err != nil {
		aroundme[mediator.South] = 2
	}

	aroundme[mediator.West], err = state.Cell(pos.X-1,pos.Y)
	if err != nil {
		aroundme[mediator.West] = 2
	}

	aroundme[mediator.East], err = state.Cell(pos.X+1,pos.Y)
	if err != nil {
		aroundme[mediator.East] = 2
	}

	return aroundme
}
