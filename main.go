package termgridboard

import (
    "github.com/briancsparks/termgridboard/internal/grid"
    "github.com/gdamore/tcell/v2"
)

func Run() {

    width, height := 40, 40

    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }
    if err := screen.Init(); err != nil {
        panic(err)
    }
    defer screen.Fini()

    layers := grid.MakeLayers(width, height, screen)
    //_ = layers

    screen.Clear()

    //quit := make(chan bool)
    chans := grid.MakeChans()

    go grid.ReadFifo(chans)

    go func() {
        for {
            select {
            case <-chans.Quit:
                return

            case cell := <- chans.Cell:
                layers.SetAt(cell)
            }
        }
    }()

    // Block, waiting for input from the user
    for {
        ev := screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
                close(chans.Quit)
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
