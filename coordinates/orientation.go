package coordinates

type Orientation uint8

const (
    North Orientation = iota
    South
    East
    West
)

func (o Orientation) String() string {
    switch o {
    case North:
        return "N"
    case South:
        return "S"
    case East:
        return "E"
    case West:
        return "W"
    }
    return "?"
}

func (o *Orientation) rotateLeft() {
    switch *o {
    case North:
        *o = West
    case West:
        *o = South
    case South:
        *o = East
    case East:
        *o = North
    }
}

func (o *Orientation) rotateRight() {
    switch *o {
    case North:
        *o = East
    case East:
        *o = South
    case South:
        *o = West
    case West:
        *o = North
    }
}
