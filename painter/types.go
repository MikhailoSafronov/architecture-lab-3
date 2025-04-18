package painter

// Texture представляет текстуру для рисования
type Texture interface{}

// OperationFunc - тип для операций рисования
type OperationFunc func(Texture)

// BgRect - прямоугольник фона
type BgRect struct {
    X1, Y1, X2, Y2 float32
}

// Move - операция перемещения
type Move struct {
    X, Y float32
}

// TFigure - фигура типа T
type TFigure struct {
    X, Y float32
    Angle int // Угол поворота для T-фигуры
}

// WhiteFill - операция заливки белым
type WhiteFill struct{}

// GreenFill - операция заливки зеленым
type GreenFill struct{}

// UpdateOp - операция обновления экрана
type UpdateOp struct{}
