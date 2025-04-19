package painter_test

import (
    "testing"
    "github.com/MikhailoSafronov/architecture-lab-3/painter"
    "github.com/stretchr/testify/assert"
)

func TestBgRectOperation(t *testing.T) {
    op := painter.BgRect{X1: 0.1, Y1: 0.2, X2: 0.3, Y2: 0.4}
    assert.Equal(t, float32(0.1), op.X1)
    // Добавьте проверки выполнения операции
}
