package main

import (
    "log"
    "github.com/roman-mazur/architecture-lab-3/ui"
    "github.com/roman-mazur/architecture-lab-3/painter"
    "github.com/roman-mazur/architecture-lab-3/painter/lang"
)

func main() {
    w, err := ui.NewWindow()
    if err != nil {
        log.Fatal(err)
    }

    loop := painter.NewLoop()
    
    // Змінити цю частину
    w.OnScreenReady = func(s screen.Screen) {
        loop.Start(s)
    }

    handler := lang.NewHandler(loop)
    handler.StartServer()

    w.Run() // Тепер викликаємо публічний метод Run
}
