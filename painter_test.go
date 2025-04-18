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
