package painter_test

import (
	"testing"
	"github.com/roman-mazur/architecture-lab-3/painter"
)

func TestBgRect(t *testing.T) {
	rect := painter.BgRect{X1: 0.1, Y1: 0.2, X2: 0.3, Y2: 0.4}
	if rect.X1 != 0.1 {
		t.Errorf("Expected X1 = 0.1, got %v", rect.X1)
	}
}
