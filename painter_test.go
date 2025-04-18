package painter

import "testing"

func TestBrushInitialization(t *testing.T) {
    b := Brush{
        Size: 10,
        Color: "black",
    }

    if b.Size != 10 || b.Color != "black" {
        t.Errorf("Expected Size=10 and Color=black, got Size=%d and Color=%s", b.Size, b.Color)
    }
}

func TestTFigureInitialization(t *testing.T) {
    f := TFigure{
        X:      5,
        Y:      10,
        Width:  100,
        Height: 50,
    }

    if f.X != 5 || f.Y != 10 || f.Width != 100 || f.Height != 50 {
        t.Errorf("TFigure init failed: got %+v", f)
    }
}

func TestMoveOperation(t *testing.T) {
    m := Move{
        DX: 3,
        DY: -2,
    }

    if m.DX != 3 || m.DY != -2 {
        t.Errorf("Expected DX=3 and DY=-2, got DX=%d and DY=%d", m.DX, m.DY)
    }
}

