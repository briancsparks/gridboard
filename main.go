package main

import (
    "github.com/gdamore/tcell/v2"
)

func main() {

    width, height := 30, 40

    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }
    if err := screen.Init(); err != nil {
        panic(err)
    }
    defer screen.Fini()

    layers := makeLayers(width, height)
    _ = layers

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
                //layers.setAt(cell)
                style := tcell.StyleDefault

                color, err := colorCode(cell.Color)
                if err == nil {
                    style = style.Foreground(color)
                }

                screen.SetCell(cell.X, cell.Y, style, cell.Sym)

                if cell.Redraw {
                    screen.Show()
                }
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
