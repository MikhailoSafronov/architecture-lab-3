package main

import "github.com/MikhailoSafronov/architecture-lab-3/lang"

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
