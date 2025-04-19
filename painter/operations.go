package painter

import "golang.org/x/exp/shiny/screen"

type Operation interface {
    Do(t screen.Texture)
}

type BgRect struct{ X1, Y1, X2, Y2 float32 }

func (op BgRect) Do(t screen.Texture) {
    // Реализация рисования прямоугольника
    t.Fill(t.Bounds(), color.RGBA{0, 0, 0, 255}, draw.Src)
}

type TFigure struct{ X, Y float32 }

func (op TFigure) Do(t screen.Texture) {
    // Реализация рисования T-фигуры
}

type Move struct{ DX, DY float32 }

func (op Move) Do(t screen.Texture) {
    // Реализация перемещения
}
