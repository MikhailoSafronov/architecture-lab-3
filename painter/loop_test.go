package painter

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
	"golang.org/x/exp/shiny/screen"
)

var size = image.Pt(800, 800)

func TestLoop_Post(t *testing.T) {
	var (
		l  Loop
		tr testReceiver
	)
	l.Receiver = &tr

	var testOps []string

	l.Start(mockScreen{})
	l.Post(BgColorOp{Color: color.White})
	l.Post(BgColorOp{Color: color.RGBA{G: 0xff, A: 0xff}})
	l.Post(UpdateOp{})

	for i := 0; i < 3; i++ {
		go l.Post(BgColorOp{Color: color.RGBA{G: 0xff, A: 0xff}})
	}

	l.Post(TestOpFunc(func(t screen.Texture, s *State) bool {
		testOps = append(testOps, "op 1")
		l.Post(TestOpFunc(func(t screen.Texture, s *State) bool {
			testOps = append(testOps, "op 2")
			return false
		}))
		return false
	}))

	l.Post(TestOpFunc(func(t screen.Texture, s *State) bool {
		testOps = append(testOps, "op 3")
		return false
	}))

	l.StopAndWait()
	// Перевірки...
}

// Перейменовано OperationFunc на TestOpFunc, щоб уникнути конфлікту
type TestOpFunc func(t screen.Texture, s *State) bool

func (f TestOpFunc) Do(t screen.Texture, s *State) bool {
	return f(t, s)
}

type testReceiver struct {
	lastTexture screen.Texture
}

func (tr *testReceiver) Update(t screen.Texture) {
	tr.lastTexture = t
}

type mockScreen struct{}

func (m mockScreen) NewBuffer(_ image.Point) (screen.Buffer, error) {
	panic("implement me")
}

func (m mockScreen) NewTexture(_ image.Point) (screen.Texture, error) {
	return new(mockTexture), nil
}

func (m mockScreen) NewWindow(_ *screen.NewWindowOptions) (screen.Window, error) {
	panic("implement me")
}

type mockTexture struct {
	Colors []color.Color
}

func (m *mockTexture) Release() {}

func (m *mockTexture) Size() image.Point { return size }

func (m *mockTexture) Bounds() image.Rectangle {
	return image.Rectangle{Max: m.Size()}
}

func (m *mockTexture) Upload(_ image.Point, _ screen.Buffer, _ image.Rectangle) {}

func (m *mockTexture) Fill(_ image.Rectangle, src color.Color, _ draw.Op) {
	m.Colors = append(m.Colors, src)
}
