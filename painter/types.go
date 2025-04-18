package painter

// Все типы в одном файле без дублирований
type Texture interface{}
type OperationFunc func(Texture)

type BgRect struct{}
type Move struct{} 
type TFigure struct{}
type WhiteFill struct{}
type GreenFill struct{}
type UpdateOp struct{}
