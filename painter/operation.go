package painter

type OperationFunc func(Texture)

type BgRect struct {
    X1, Y1, X2, Y2 float32
}

type Move struct {
    X, Y float32
}
