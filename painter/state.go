package painter

import (
	"image"
	"image/color"
)

type State struct {
	bgColor   color.Color
	bgRect    image.Rectangle
	figures   []image.Rectangle
	moveDelta image.Point
}

func (s *State) Reset() {
	s.bgColor = color.Black
	s.bgRect = image.Rect(0, 0, 0, 0)
	s.figures = nil
	s.moveDelta = image.Point{}
}
