package coordinates

import "testing"

func TestOrientation(t *testing.T) {
    o := North

    // tests moving counter-clock wise
    o.rotateLeft()
    if o != West {
        t.Errorf("Orientation should be West: %v", o)
    }

    o.rotateLeft()
    if o != South {
        t.Errorf("Orientation should be South: %v", o)
    }

    o.rotateLeft()
    if o != East {
        t.Errorf("Orientation should be East: %v", o)
    }

    o.rotateLeft()
    if o != North {
        t.Errorf("Orientation should be North: %v", o)
    }

    // tests moving clock wise
    o.rotateRight()
    if o != East {
        t.Errorf("Orientation should be East: %v", o)
    }

    o.rotateRight()
    if o != South {
        t.Errorf("Orientation should be South: %v", o)
    }

    o.rotateRight()
    if o != West {
        t.Errorf("Orientation should be West: %v", o)
    }

    o.rotateRight()
    if o != North {
        t.Errorf("Orientation should be North: %v", o)
    }

    // Orientation string
    if North.String() != "N" {
        t.Errorf("West should be: 'N': %s", West)
    }
    if South.String() != "S" {
        t.Errorf("South should be: 'S': %s", South)
    }
    if West.String() != "W" {
        t.Errorf("West should be: 'W': %s", West)
    }
    if East.String() != "E" {
        t.Errorf("East should be: 'E': %s", East)
    }
    unkown := East + 6
    if unkown.String() != "?" {
        t.Errorf("Unknown should be: '?': %s", unkown)
    }
}
