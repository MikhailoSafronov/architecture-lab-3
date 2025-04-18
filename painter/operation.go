package painter

type Texture interface{}

type OperationFunc func(Texture)

type BgRect struct { X1, Y1, X2, Y2 float32 }
type Move struct { X, Y float32 }
type TFigure struct { X, Y float32 }
type WhiteFill struct{}
