package painter

func (op OperationFunc) Draw(t Texture) {
    op(t)
}

func (r BgRect) Draw(t Texture) {
    // реализация
}

func (m Move) Draw(t Texture) {
    // реализация
}

func (f TFigure) Draw(t Texture) {
    // реализация
}
