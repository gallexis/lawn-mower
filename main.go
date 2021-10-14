package main

import (
    "fmt"
    "mow/coordinates"
    "os"
    "sync"
)

func main() {
    args := os.Args
    if len(args) < 2 {
        panic("The path of the input file is required")
    }

    filepath := args[1]

    // Read and parse the input file
    lawnCoordinates, mowersCoordinates, err := ParseFile(filepath)
    if err != nil {
        panic(err)
    }

    // Load the Lawn & Mowers from the parsed file
    startChan := make(chan struct{})
    mowers := getMowers(mowersCoordinates, startChan)
    lawn := NewLawn(lawnCoordinates.X, lawnCoordinates.Y, mowers)

    wg := sync.WaitGroup{}
    lawn.Run(&wg)

    for _, mower := range mowers {
        wg.Add(len(mower.Directions))
        go mower.Start()
    }

    // Broadcast a start signal to all Mowers
    close(startChan)

    // Wait for all Mowers to finish
    wg.Wait()

    // Display the result
    for _, m := range mowers {
        fmt.Printf("%s\n", m)
    }
}

func getMowers(mowersCoordinates []MowerCoordinates, startChan chan struct{}) []*Mower {
    var mowers = make([]*Mower, len(mowersCoordinates))

    for i, mower := range mowersCoordinates {
        mowerPosition := coordinates.Position{
            X: mower.X,
            Y: mower.Y,
            O: mower.O,
        }

        mowers[i] = NewMower(mowerPosition, mower.Directions, startChan)
    }

    return mowers
}
