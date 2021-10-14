package coordinates

import "fmt"

type Position struct {
    X int
    Y int
    O Orientation
}

func (p Position) Hash() string {
    return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func (p Position) NextPosition(d Direction) Position {
    switch d {
    case L:
        p.O.rotateLeft()

    case R:
        p.O.rotateRight()

    case F:
        switch p.O {
        case North:
            p.Y++
        case South:
            p.Y--
        case East:
            p.X++
        case West:
            p.X--
        }
    }

    return p
}
