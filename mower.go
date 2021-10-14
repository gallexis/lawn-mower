package main

import (
    "fmt"
    "mow/coordinates"
)

type Mower struct {
    CurrentPosition     coordinates.Position
    Directions          []coordinates.Direction
    processedDirections int

    startChan chan struct{}
    lawnChan  chan *Mower
}

func NewMower(initPosition coordinates.Position, directions []coordinates.Direction, startChan chan struct{}) *Mower {
    return &Mower{
        CurrentPosition:     initPosition,
        Directions:          directions,
        processedDirections: 0,
        startChan:           startChan,
    }
}

func (m Mower) String() string {
    x := m.CurrentPosition.X
    y := m.CurrentPosition.Y
    o := m.CurrentPosition.O

    return fmt.Sprintf("%d %d %s", x, y, o)
}

func (m Mower) CalculateNextPosition() coordinates.Position {
    direction := m.Directions[m.processedDirections]
    return m.CurrentPosition.NextPosition(direction)
}

func (m *Mower) Move(nextPosition coordinates.Position) {
    m.CurrentPosition = nextPosition
    m.processedDirections++
}

func (m *Mower) DiscardMove() {
    m.processedDirections++
}

/*
	Start waits for a broadcast "start" signal,
	then sends itself to the lawnChan to be processed by the Lawn
*/
func (m *Mower) Start() {
    <-m.startChan

    for range m.Directions {
        m.lawnChan <- m
    }
}

func (m *Mower) AddLawnChan(lawnChan chan *Mower) {
    m.lawnChan = lawnChan
}
