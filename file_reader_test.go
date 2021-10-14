package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"syscall"
	"testing"

	c "mow/coordinates"
)

func TestParseFile(t *testing.T) {
	file, err := os.CreateTemp("", "data")
	if err != nil {
		t.Error(err)
	}
	defer syscall.Unlink(file.Name())
	ioutil.WriteFile(file.Name(), []byte(`10 10
1 2 N
LFLFLFLFF
3 3 E
FFRFFRFRRF
4 7 S
FFL
`), 0644)

	lawnCoordinates, mowersCoordinates, err := ParseFile(file.Name())
	if err != nil {
		t.Error(err)
	}

	if lawnCoordinates.X != 10 || lawnCoordinates.Y != 10 {
		t.Errorf("wrong lawn coordinates %v", lawnCoordinates)
	}

	if mowersCoordinates[0].X != 1 || mowersCoordinates[0].Y != 2 || mowersCoordinates[0].O != c.North {
		t.Errorf("wrong mower coordinates %v", mowersCoordinates[0])
	}
	if !reflect.DeepEqual(mowersCoordinates[0].Directions,
		[]c.Direction{c.L, c.F, c.L, c.F, c.L, c.F, c.L, c.F, c.F}) {
		t.Errorf("wrong mower directions %v", mowersCoordinates[0])
	}

	if mowersCoordinates[1].X != 3 || mowersCoordinates[1].Y != 3 || mowersCoordinates[1].O != c.East {
		t.Errorf("wrong mower coordinates %v", mowersCoordinates[1])
	}
	if !reflect.DeepEqual(mowersCoordinates[1].Directions,
		[]c.Direction{c.F, c.F, c.R, c.F, c.F, c.R, c.F, c.R, c.R, c.F}) {
		t.Errorf("wrong mower directions %v", mowersCoordinates[1])
	}

	if mowersCoordinates[2].X != 4 || mowersCoordinates[2].Y != 7 || mowersCoordinates[2].O != c.South {
		t.Errorf("wrong mower coordinates %v", mowersCoordinates[2])
	}
	if !reflect.DeepEqual(mowersCoordinates[2].Directions,
		[]c.Direction{c.F, c.F, c.L}) {
		t.Errorf("wrong mower directions %v", mowersCoordinates[2])
	}

}
