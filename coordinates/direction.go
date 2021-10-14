package coordinates

type Direction uint8

const (
    L Direction = iota
    R
    F
)

func (d Direction) String() string {
    switch d {
    case L:
        return "L"
    case R:
        return "R"
    case F:
        return "F"
    }
    return "?"
}
