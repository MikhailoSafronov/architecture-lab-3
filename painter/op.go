package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

type Operation interface {
	Do(t screen.Texture, s *State) bool
}
type ResetOp struct{}

func (op ResetOp) Do(t screen.Texture, s *State) bool {
	s.Reset()
	t.Fill(t.Bounds(), color.Black, screen.Src)
	return true
}

type OperationList []Operation

func (ol OperationList) Do(t screen.Texture, s *State) bool {
	var updated bool
	for _, op := range ol {
		if op.Do(t, s) {
			updated = true
		}
	}
	return updated
}

type BgColorOp struct {
	Color color.Color
}

func (op BgColorOp) Do(t screen.Texture, s *State) bool {
	s.bgColor = op.Color
	t.Fill(t.Bounds(), op.Color, screen.Src)
	return false
}

type BgRectOp struct {
	X1, Y1, X2, Y2 float64
}

func (op BgRectOp) Do(t screen.Texture, s *State) bool {
	x1 := int(op.X1 * 800)
	y1 := int(op.Y1 * 800)
	x2 := int(op.X2 * 800)
	y2 := int(op.Y2 * 800)
	s.bgRect = image.Rect(x1, y1, x2, y2)
	t.Fill(s.bgRect, color.Black, screen.Src)
	return false
}

type FigureOp struct {
	X, Y float64
}

func (op FigureOp) Do(t screen.Texture, s *State) bool {
	size := 100
	x := int(op.X*800) - size/2 + s.moveDelta.X
	y := int(op.Y*800) - size/2 + s.moveDelta.Y

	// Малювання T-фігури
	t.Fill(image.Rect(x+20, y, x+80, y+60), color.RGBA{R: 255, G: 255, B: 0, A: 255}, screen.Src)
	t.Fill(image.Rect(x, y+40, x+100, y+60), color.RGBA{R: 255, G: 255, B: 0, A: 255}, screen.Src)

	s.figures = append(s.figures, image.Rect(x, y, x+100, y+100))
	return false
}

type MoveOp struct {
	X, Y float64
}

func (op MoveOp) Do(_ screen.Texture, s *State) bool {
	s.moveDelta = image.Pt(int(op.X*800), int(op.Y*800))
	return false
}

type UpdateOp struct{}

func (op UpdateOp) Do(_ screen.Texture, _ *State) bool {
	return true
}
