package painter

import "golang.org/x/exp/shiny/screen"

type Receiver interface {
	Update(t screen.Texture)
}
