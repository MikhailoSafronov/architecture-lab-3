package painter_test

import (
	"strings"
	"testing"
	"github.com/roman-mazur/architecture-lab-3/painter"
	"github.com/roman-mazur/architecture-lab-3/painter/lang"
	"github.com/stretchr/testify/assert"
)


func TestLoop(t *testing.T) {
	t.Run("Post and receive operation", func(t *testing.T) {
		var l painter.Loop
		called := false
		
		l.Post(painter.OperationFunc(func(painter.Texture) {
			called = true
		}))
		
		l.Start()
		defer l.Stop()
		
		assert.Eventually(t, func() bool { return called }, time.Second, 10*time.Millisecond)
	})
}


func TestParser(t *testing.T) {
	p := lang.Parser{}
	
	tests := []struct {
		name    string
		input   string
		wantErr bool
		wantOp  painter.Operation
	}{
		{
			name:    "valid bgrect",
			input:   "bgrect 0.1 0.1 0.5 0.5",
			wantOp:  painter.BgRect{X1: 0.1, Y1: 0.1, X2: 0.5, Y2: 0.5},
		},
		{
			name:    "invalid bgrect",
			input:   "bgrect invalid",
			wantErr: true,
		},
		{
			name:    "move command",
			input:   "move 0.2 0.3",
			wantOp:  painter.Move{X: 0.2, Y: 0.3},
		},
		{
			name:    "white command",
			input:   "white",
			wantOp:  painter.WhiteFill{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ops, err := p.Parse(strings.NewReader(tt.input))
			
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			
			assert.NoError(t, err)
			assert.IsType(t, tt.wantOp, ops[0])
		})
	}
}


func TestOperations(t *testing.T) {
	t.Run("BgRect coordinates", func(t *testing.T) {
		rect := painter.BgRect{X1: 0.1, Y1: 0.2, X2: 0.3, Y2: 0.4}
		assert.Equal(t, float32(0.1), rect.X1)
		assert.Equal(t, float32(0.4), rect.Y2)
	})

	t.Run("Move operation", func(t *testing.T) {
		move := painter.Move{X: 0.5, Y: 0.6}
		assert.Equal(t, float32(0.5), move.X)
	})

	t.Run("TFigure position", func(t *testing.T) {
		fig := painter.TFigure{X: 0.7, Y: 0.8}
		assert.Equal(t, float32(0.7), fig.X)
	})
}
