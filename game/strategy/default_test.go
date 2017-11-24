package strategy_test

import (
	"testing"

	"mktd5/mktd-island/client/game/strategy"
	"mktd5/mktd-island/client/mediator"
)

func TestWhatsAroundMe(t *testing.T) {
	sampleMap := [][]mediator.Cell{
		{0,1,0},
		{3,0,2},
		{0,0,0},
	}


	result := strategy.WhatsAroundMe(1,1, sampleMap)
	if result[mediator.North] != 1 {
		t.Errorf("North should be 1, got %v", result[mediator.North])
	}
	if result[mediator.East] != 2 {
		t.Errorf("East should be 2, got %v", result[mediator.East])
	}
	if result[mediator.South] != 0 {
		t.Errorf("South should be 0, got %v", result[mediator.South])
	}
	if result[mediator.West] != 3 {
		t.Errorf("West should be 3, got %v", result[mediator.West])
	}

}

func TestWhatsAroundMe1(t *testing.T) {
	sampleMap := [][]mediator.Cell{
		{0,1,0},
		{3,0,2},
		{0,0,0},
	}

	result := strategy.WhatsAroundMe(0,0, sampleMap)
	if result[mediator.North] != 2 {
		t.Errorf("North should be 2, got %v", result[mediator.North])
	}
	if result[mediator.East] != 1 {
		t.Errorf("East should be 1, got %v", result[mediator.East])
	}
	if result[mediator.South] != 3 {
		t.Errorf("South should be 3, got %v", result[mediator.South])
	}
	if result[mediator.West] != 2 {
		t.Errorf("West should be 2, got %v", result[mediator.West])
	}
}

func TestWhatsAroundMe2(t *testing.T) {
	sampleMap := [][]mediator.Cell{
		{0,1,0},
		{3,0,2},
		{0,0,0},
	}


	result := strategy.WhatsAroundMe(3,3, sampleMap)
	if result[mediator.North] != 2 {
		t.Errorf("North should be 2, got %v", result[mediator.North])
	}
	if result[mediator.East] != 2 {
		t.Errorf("East should be 2, got %v", result[mediator.East])
	}
	if result[mediator.South] != 2 {
		t.Errorf("South should be 2, got %v", result[mediator.South])
	}
	if result[mediator.West] != 0 {
		t.Errorf("West should be 0, got %v", result[mediator.West])
	}
}