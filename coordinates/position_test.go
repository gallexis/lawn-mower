package coordinates

import "testing"

func TestPosition(t *testing.T) {
    p := Position{
        X: 2,
        Y: 2,
        O: North,
    }

    p2 := p.NextPosition(F)
    expectedP2 := Position{
        X: 2,
        Y: 3,
        O: North,
    }
    if p2 != expectedP2 {
        t.Errorf("Expected p2: %+v, got: %+v", expectedP2, p2)
    }

    p3 := p.NextPosition(L)
    expectedP3 := Position{
        X: 2,
        Y: 2,
        O: West,
    }
    if p3 != expectedP3 {
        t.Errorf("Expected p3: %+v, got: %+v", expectedP3, p3)
    }

    p.O = South
    p4 := p.NextPosition(F)
    expectedP4 := Position{
        X: 2,
        Y: 1,
        O: South,
    }
    if p4 != expectedP4 {
        t.Errorf("Expected p4: %+v, got: %+v", expectedP4, p4)
    }

    p.O = West
    p5 := p.NextPosition(F)
    expectedP5 := Position{
        X: 1,
        Y: 2,
        O: West,
    }
    if p5 != expectedP5 {
        t.Errorf("Expected p5: %+v, got: %+v", expectedP5, p5)
    }

    p.O = East
    p6 := p.NextPosition(F)
    expectedP6 := Position{
        X: 3,
        Y: 2,
        O: East,
    }
    if p6 != expectedP6 {
        t.Errorf("Expected p6: %+v, got: %+v", expectedP6, p6)
    }

    p7 := p.NextPosition(R)
    expectedP7 := Position{
        X: 2,
        Y: 2,
        O: South,
    }
    if p7 != expectedP7 {
        t.Errorf("Expected p7: %+v, got: %+v", expectedP7, p7)
    }

    // Position hash
    if p7.Hash() != "2-2" {
        t.Errorf("p7 should be: '2-2': %s", p7.Hash())
    }
}
