package painter_test

import (
	"testing"
	"github.com/roman-mazur/architecture-lab-3/painter"
	"github.com/roman-mazur/architecture-lab-3/painter/lang"
	"github.com/stretchr/testify/assert"
)

func TestLoop_Post(t *testing.T) {
	var l painter.Loop
	l.Post(painter.OperationFunc(func(painter.Texture) {
		// Тестовая операция
	}))
	assert.NotNil(t, l, "Loop should be initialized")
}

func TestParser_ValidBgRect(t *testing.T) {
	p := lang.Parser{}
	ops, err := p.Parse("bgrect 0.1 0.1 0.5 0.5")
	assert.NoError(t, err)
	assert.IsType(t, painter.BgRect{}, ops[0], "Should return BgRect operation")
}

func TestParser_InvalidBgRect(t *testing.T) {
	p := lang.Parser{}
	_, err := p.Parse("bgrect invalid")
	assert.Error(t, err, "Should return error for invalid command")
}

func TestMoveOperation(t *testing.T) {
	moveOp := painter.Move{OffsetX: 0.5, OffsetY: 0.5}
	assert.Equal(t, 0.5, moveOp.OffsetX, "Move operation should store X offset")
	assert.Equal(t, 0.5, moveOp.OffsetY, "Move operation should store Y offset")
}

func TestTFigure(t *testing.T) {
	figure := painter.TFigure{CenterX: 0.5, CenterY: 0.5}
	assert.Equal(t, 0.5, figure.CenterX, "TFigure should store X coordinate")
	assert.Equal(t, 0.5, figure.CenterY, "TFigure should store Y coordinate")
}
