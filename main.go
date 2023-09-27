package main

import (
    "github.com/gdamore/tcell/v2"
)

func main() {

    width, height := 40, 40

    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }
    if err := screen.Init(); err != nil {
        panic(err)
    }
    defer screen.Fini()

    layers := makeLayers(width, height, screen)
    //_ = layers

    screen.Clear()

    //quit := make(chan bool)
    chans := makeChans()

    go readFifo(chans)

    go func() {
        for {
            select {
            case <-chans.quit:
                return

            case cell := <- chans.cell:
                layers.setAt(cell)
            }
        }
    }()

    // Block, waiting for input from the user
    for {
        ev := screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
                close(chans.quit)
                return
            }
            if ev.Key() == tcell.KeyCtrlL {
                screen.Sync()
            }
        case *tcell.EventResize:
            screen.Sync()
        }
    }

}
