package main

import (
    "bufio"
    "errors"
    "fmt"
    "mow/coordinates"
    "os"
    "strconv"
    "strings"
)

type LawnCoordinates struct {
    X int
    Y int
}

func (c *LawnCoordinates) parseFromString(lawnCoordinates string) error {
    coordinatesSlice := strings.Split(lawnCoordinates, " ")
    if len(coordinatesSlice) != 2 {
        return errors.New("incorrect coordinates")
    }

    x, err := strconv.ParseInt(coordinatesSlice[0], 10, 64)
    if err != nil {
        return err
    }

    y, err := strconv.ParseInt(coordinatesSlice[1], 10, 64)
    if err != nil {
        return err
    }

    c.X = int(x)
    c.Y = int(y)

    return nil
}

type MowerCoordinates struct {
    X          int
    Y          int
    O          coordinates.Orientation
    Directions []coordinates.Direction
}

func (c *MowerCoordinates) parseCoordinates(mowerCoordinates string) error {
    coordinatesSlice := strings.Split(mowerCoordinates, " ")
    if len(coordinatesSlice) != 3 {
        return fmt.Errorf("incorrect mower coordinates: %s", mowerCoordinates)
    }

    x, err := strconv.ParseInt(coordinatesSlice[0], 10, 64)
    if err != nil {
        return err
    }
    c.X = int(x)

    y, err := strconv.ParseInt(coordinatesSlice[1], 10, 64)
    if err != nil {
        return err
    }
    c.Y = int(y)

    orientation := coordinatesSlice[2]
    switch orientation {
    case "N":
        c.O = coordinates.North
    case "S":
        c.O = coordinates.South
    case "W":
        c.O = coordinates.West
    case "E":
        c.O = coordinates.East
    default:
        return fmt.Errorf("wrong orientation: %s", orientation)
    }

    return nil
}

func (c *MowerCoordinates) parseDirections(directions string) error {
    for _, dir := range directions {
        direction := string(dir)
        switch direction {
        case "L":
            c.Directions = append(c.Directions, coordinates.L)
        case "R":
            c.Directions = append(c.Directions, coordinates.R)
        case "F":
            c.Directions = append(c.Directions, coordinates.F)
        default:
            return fmt.Errorf("wrong direction: %s", direction)
        }
    }

    return nil
}

func ParseFile(path string) (lawnCoordinates LawnCoordinates, mowersCoordinates []MowerCoordinates, err error) {
    file, err := os.Open(path)
    if err != nil {
        return
    }
    defer file.Close()

    var line string
    scanner := bufio.NewScanner(bufio.NewReader(file))

    // read the Lawn dimensions
    scanner.Scan()
    line = scanner.Text()

    err = lawnCoordinates.parseFromString(line)
    if err != nil {
        return
    }

    // read the Mowers coordinates & directions
    for scanner.Scan() {
        mower := MowerCoordinates{}

        err = mower.parseCoordinates(scanner.Text())
        if err != nil {
            return
        }

        scanner.Scan()
        err = mower.parseDirections(scanner.Text())
        if err != nil {
            return
        }

        mowersCoordinates = append(mowersCoordinates, mower)
    }

    return
}
