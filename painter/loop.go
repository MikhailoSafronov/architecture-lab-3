package painter

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
)

type Loop struct {
	Receiver Receiver
	screen   screen.Screen
	next     screen.Texture
	prev     screen.Texture
	state    State
	mq       chan Operation
	stop     chan struct{}
	stopOnce sync.Once
}

func (l *Loop) Start(s screen.Screen) {
	l.screen = s
	l.mq = make(chan Operation, 100)
	l.stop = make(chan struct{})
	l.state.Reset()

	var err error
	l.next, err = l.screen.NewTexture(image.Pt(800, 800))
	if err != nil {
		panic(err)
	}

	go l.eventLoop()
}

func (l *Loop) eventLoop() {
	for {
		select {
		case op := <-l.mq:
			if op.Do(l.next, &l.state) {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next

				newTex, _ := l.screen.NewTexture(image.Pt(800, 800))
				l.next = newTex
			}
		case <-l.stop:
			return
		}
	}
}

func (l *Loop) Post(op Operation) {
	l.mq <- op
}

func (l *Loop) StopAndWait() {
	l.stopOnce.Do(func() {
		close(l.stop)
	})
}
