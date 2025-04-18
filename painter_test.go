package painter

import "testing"

func TestTypesExist(t *testing.T) {

    var _ Texture
    var _ OperationFunc
    var _ = BgRect{}
    var _ = Move{}
    var _ = TFigure{}
    var _ = WhiteFill{}
    var _ = GreenFill{}
    var _ = UpdateOp{}
}
