package main

import (
	"sync"
	"testing"

	"mow/coordinates"
)

func TestIsValidNextPosition(t *testing.T) {
	startChan := make(chan struct{})

	// test valid move
	mower1 := NewMower(coordinates.Position{X: 2, Y: 2, O: coordinates.North}, []coordinates.Direction{coordinates.F}, startChan)
	l := NewLawn(5, 5, []*Mower{mower1})
	valid := l.isValidNextPosition(mower1.CurrentPosition, mower1.CalculateNextPosition())
	if !valid {
		t.Error("should be valid move")
	}

	// test outbound moves
	mower1 = NewMower(coordinates.Position{X: 5, Y: 5, O: coordinates.North}, []coordinates.Direction{coordinates.F}, startChan)
	l = NewLawn(5, 5, []*Mower{mower1})
	valid = l.isValidNextPosition(mower1.CurrentPosition, mower1.CalculateNextPosition())
	if valid {
		t.Error("should be invalid move")
	}

	mower1 = NewMower(coordinates.Position{X: 5, Y: 5, O: coordinates.East}, []coordinates.Direction{coordinates.F}, startChan)
	l = NewLawn(5, 5, []*Mower{mower1})
	valid = l.isValidNextPosition(mower1.CurrentPosition, mower1.CalculateNextPosition())
	if valid {
		t.Error("should be invalid move")
	}

	mower1 = NewMower(coordinates.Position{X: 5, Y: 5, O: coordinates.South}, []coordinates.Direction{coordinates.F}, startChan)
	l = NewLawn(0, 0, []*Mower{mower1})
	valid = l.isValidNextPosition(mower1.CurrentPosition, mower1.CalculateNextPosition())
	if valid {
		t.Error("should be invalid move")
	}

	mower1 = NewMower(coordinates.Position{X: 0, Y: 0, O: coordinates.West}, []coordinates.Direction{coordinates.F}, startChan)
	l = NewLawn(5, 5, []*Mower{mower1})
	valid = l.isValidNextPosition(mower1.CurrentPosition, mower1.CalculateNextPosition())
	if valid {
		t.Error("should be invalid move")
	}

	// test collision
	mower1 = NewMower(coordinates.Position{X: 0, Y: 0, O: coordinates.East}, []coordinates.Direction{coordinates.F}, startChan)
	mower2 := NewMower(coordinates.Position{X: 1, Y: 0, O: coordinates.West}, []coordinates.Direction{coordinates.F}, startChan)
	l = NewLawn(5, 5, []*Mower{mower1, mower2})
	valid = l.isValidNextPosition(mower1.CurrentPosition, mower1.CalculateNextPosition())
	if valid {
		t.Error("should be invalid move")
	}

	// test only changing direction
	mower1 = NewMower(coordinates.Position{X: 1, Y: 0, O: coordinates.West}, []coordinates.Direction{coordinates.L}, startChan)
	l = NewLawn(5, 5, []*Mower{mower1})
	valid = l.isValidNextPosition(mower1.CurrentPosition, mower1.CalculateNextPosition())
	if !valid {
		t.Error("should be valid move")
	}
}

func TestRun(t *testing.T) {
	// Moving mowers
	startChan := make(chan struct{})
	mower1 := NewMower(coordinates.Position{X: 5, Y: 5, O: coordinates.East}, []coordinates.Direction{coordinates.F}, startChan)
	mower2 := NewMower(coordinates.Position{X: 1, Y: 0, O: coordinates.West}, []coordinates.Direction{coordinates.F}, startChan)

	mowers := []*Mower{mower1, mower2}
	lawn := NewLawn(10, 10, mowers)

	wg := sync.WaitGroup{}
	lawn.Run(&wg)

	for _, m := range mowers {
		wg.Add(len(m.Directions))
		go m.Start()
	}

	close(startChan)
	wg.Wait()

	if mower1.CurrentPosition.X != 6 || mower1.CurrentPosition.Y != 5 || mower1.CurrentPosition.O != coordinates.East {
		t.Errorf("incorrect mower1 position: %+v", mower1)
	}
	if mower2.CurrentPosition.X != 0 || mower2.CurrentPosition.Y != 0 || mower2.CurrentPosition.O != coordinates.West {
		t.Errorf("incorrect mower2 position: %+v", mower2)
	}

	// static mowers due to collision
	startChan = make(chan struct{})
	mower1 = NewMower(coordinates.Position{X: 0, Y: 0, O: coordinates.East}, []coordinates.Direction{coordinates.F}, startChan)
	mower2 = NewMower(coordinates.Position{X: 1, Y: 0, O: coordinates.West}, []coordinates.Direction{coordinates.F}, startChan)

	mowers = []*Mower{mower1, mower2}
	lawn = NewLawn(10, 10, mowers)

	wg = sync.WaitGroup{}
	lawn.Run(&wg)

	for _, m := range mowers {
		wg.Add(len(m.Directions))
		go m.Start()
	}

	close(startChan)
	wg.Wait()

	if mower1.CurrentPosition.X != 0 || mower1.CurrentPosition.Y != 0 || mower1.CurrentPosition.O != coordinates.East {
		t.Errorf("incorrect mower1 position: %+v", mower1)
	}
	if mower2.CurrentPosition.X != 1 || mower2.CurrentPosition.Y != 0 || mower2.CurrentPosition.O != coordinates.West {
		t.Errorf("incorrect mower2 position: %+v", mower2)
	}
}
