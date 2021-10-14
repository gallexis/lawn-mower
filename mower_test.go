package main

import (
	"testing"

	"mow/coordinates"
)

func TestMower(t *testing.T) {
	initPosition := coordinates.Position{
		X: 4,
		Y: 4,
		O: coordinates.North,
	}
	directions := []coordinates.Direction{coordinates.F, coordinates.L, coordinates.F, coordinates.L, coordinates.F, coordinates.R, coordinates.F}
	mower := NewMower(initPosition, directions, make(chan struct{}))

	for range directions {
		nextPosition := mower.CalculateNextPosition()
		mower.Move(nextPosition)
	}

	finalPosition := coordinates.Position{
		X: 2,
		Y: 4,
		O: coordinates.West,
	}

	if mower.CurrentPosition != finalPosition {
		t.Errorf("Expected final position to be: %+v, got: %+v", finalPosition, mower.CurrentPosition)
	}

	// Mower String
	if mower.String() != "2 4 W" {
		t.Errorf("Mower string should be: 2 4 W, got: %s", mower)
	}
}
