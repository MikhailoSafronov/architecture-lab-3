package painter_test

import (
    "testing"
    "github.com/roman-mazur/architecture-lab-3/painter"
    "github.com/stretchr/testify/assert"
)

func TestBgRect(t *testing.T) {
    rect := painter.BgRect{X1: 0.1, Y1: 0.2, X2: 0.3, Y2: 0.4}
    assert.Equal(t, float32(0.1), rect.X1)
    assert.Equal(t, float32(0.4), rect.Y2)
}

func TestTFigure(t *testing.T) {
    fig := painter.TFigure{X: 0.5, Y: 0.6, Angle: 90}
    assert.Equal(t, float32(0.5), fig.X)
    assert.Equal(t, 90, fig.Angle)
}

func TestMove(t *testing.T) {
    move := painter.Move{X: 0.7, Y: 0.8}
    assert.Equal(t, float32(0.7), move.X)
}
