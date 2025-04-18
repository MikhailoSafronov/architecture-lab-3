package painter

// Реализация операций
func (op OperationFunc) Draw(t Texture) {
    op(t)
}

func (r BgRect) Draw(t Texture) {
    // Логика рисования прямоугольника
}

func (m Move) Draw(t Texture) {
    // Логика перемещения
}

func (f TFigure) Draw(t Texture) {
    // Логика рисования T-фигуры
}
