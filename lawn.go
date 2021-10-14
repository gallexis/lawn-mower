package main

import (
	"fmt"
	"sync"

	"mow/coordinates"
)

type Lawn struct {
	XMax           int
	YMax           int
	lawnChan       chan *Mower
	mowersPosition map[string]struct{}
}

func NewLawn(xmax, ymax int, mowers []*Mower) *Lawn {
	l := &Lawn{
		XMax:           xmax,
		YMax:           ymax,
		mowersPosition: make(map[string]struct{}),
		lawnChan:       make(chan *Mower),
	}

	for _, mower := range mowers {
		mower.AddLawnChan(l.lawnChan)

		currentPositionHash := mower.CurrentPosition.Hash()
		if _, ok := l.mowersPosition[currentPositionHash]; ok {
			panic("mowers already present at this position")
		}

		l.mowersPosition[currentPositionHash] = struct{}{}
	}

	return l
}

/*
	Run runs a goroutine that will wait for new mowers to be received on the lawnChan channel,
	then calculate if their next position is valid.
	If yes, it will move it to the next position, then update the Lawn's mowersPosition map
	If not, it will discard the mower's move
*/
func (l *Lawn) Run(wg *sync.WaitGroup) {
	go func() {
		for {
			mower := <-l.lawnChan
			currentPosition := mower.CurrentPosition
			nextPosition := mower.CalculateNextPosition()

			if l.isValidNextPosition(currentPosition, nextPosition) {
				l.updateMowersPositionMap(currentPosition, nextPosition)
				mower.Move(nextPosition)
			} else {
				fmt.Println("Invalid next position")
				mower.DiscardMove()
			}

			wg.Done()
		}
	}()
}

func (l *Lawn) isValidNextPosition(currentPosition, nextPosition coordinates.Position) bool {
	// ignore if only orientation changed
	if currentPosition.X == nextPosition.X && currentPosition.Y == nextPosition.Y {
		return true
	}

	if nextPosition.Y < 0 || nextPosition.Y > l.YMax {
		return false
	}
	if nextPosition.X < 0 || nextPosition.X > l.XMax {
		return false
	}

	// we check if there is no other mower in the next position
	if _, exist := l.mowersPosition[nextPosition.Hash()]; exist {
		return false
	}

	return true
}

func (l *Lawn) updateMowersPositionMap(oldPosition, newPosition coordinates.Position) {
	delete(l.mowersPosition, oldPosition.Hash())
	l.mowersPosition[newPosition.Hash()] = struct{}{}
}
