package coordinates

import "testing"

func TestDirection(t *testing.T) {
    // Direction string
    if L.String() != "L" {
        t.Errorf("West should be: 'L': %s", L)
    }
    if R.String() != "R" {
        t.Errorf("South should be: 'R': %s", R)
    }
    if F.String() != "F" {
        t.Errorf("West should be: 'F': %s", F)
    }
    unkown := F + 6
    if unkown.String() != "?" {
        t.Errorf("Unknown should be: '?': %s", unkown)
    }
}
