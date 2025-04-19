package main

import (
    "log"
    "github.com/roman-mazur/architecture-lab-3/ui"
    "github.com/roman-mazur/architecture-lab-3/painter"
    "github.com/roman-mazur/architecture-lab-3/lang"
)

func main() {
    w, err := ui.NewWindow()
    if err != nil {
        log.Fatal(err)
    }

    loop := painter.NewLoop()
    loop.Start(w)

    handler := lang.NewHandler(loop)
    handler.StartServer()

    w.Run()
}
