package painter

import "testing"

func TestTypesExist(t *testing.T) {
    // Просто проверяем что типы доступны
    var _ Texture = nil
    var _ OperationFunc = nil
    var _ = BgRect{}
    var _ = Move{}
    var _ = TFigure{} 
    var _ = WhiteFill{}
    var _ = GreenFill{}
    var _ = UpdateOp{}
}
