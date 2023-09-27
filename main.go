package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "github.com/gdamore/tcell/v2"
    "os"
)

type Cell struct {
    X      int    `json:"x"`
    Y      int    `json:"y"`
    Sym    rune   `json:"sym"`
    Color  string `json:"color"`
    Redraw bool   `json:"redraw"`
}

type Chans struct {
    cell chan Cell
}

func makeChans() *Chans {
    var c Chans

    c.cell = make(chan Cell, 100)

    return &c
}

func readFifo(line chan string, chans *Chans) {
    for {
        f, err := os.OpenFile("/tmp/termgridboard", os.O_RDONLY, 0600)
        if err != nil {
            panic(err)
        }

        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            text := scanner.Text()

            var cell Cell
            if err := json.Unmarshal([]byte(text), &cell); err != nil {
                continue // Skip invalid JSON
            }

            //line <- text
            chans.cell <- cell
        }
        f.Close()
    }
}

func main() {
    fmt.Println("Hello, World!")

    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }
    if err := screen.Init(); err != nil {
        panic(err)
    }
    defer screen.Fini()

    screen.Clear()

    f, err := os.OpenFile("/tmp/termgridboard", os.O_RDONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    quit := make(chan bool)
    line := make(chan string, 100)
    chans := makeChans()

    go readFifo(line, chans)

    problemJson := ""

    go func() {
        for {
            select {
            case <-quit:
                return

            case l := <-line:
                var cell Cell
                if err := json.Unmarshal([]byte(l), &cell); err != nil {
                   problemJson = l
                   continue // Skip invalid JSON
                }

                if cell.X >= 0 && cell.Y >= 0 {
                    style := tcell.StyleDefault

                    color, err := colorCode(cell.Color)
                    if err == nil {
                        style = style.Foreground(color)
                    }

                    screen.SetCell(cell.X, cell.Y, style, rune(cell.Sym))
                }

                if cell.Redraw {
                    screen.Show()
                }

            case cell := <- chans.cell:
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
                fmt.Print("Problem JSON: ")
                fmt.Println(problemJson)
                close(quit)
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

func colorCode(name string) (tcell.Color, error) {
    if name == "red" {
        return tcell.ColorRed, nil
    }
    if name == "teal" {
        return tcell.ColorTeal, nil
    }
    if name == "white" {
        return tcell.ColorWhite, nil
    }
    if name == "fuchsia" {
        return tcell.ColorFuchsia, nil
    }
    if name == "hotpink" {
        return tcell.ColorFuchsia, nil
    }
    if name == "yellow" {
        return tcell.ColorYellow, nil
    }
    if name == "gray" {
        return tcell.ColorGray, nil
    }

    return tcell.ColorWhite, fmt.Errorf("unknown color: %v", name)
}

// Colors:
//
// ColorBlack
// ColorMaroon
// ColorGreen
// ColorOlive
// ColorNavy
// ColorPurple
// ColorTeal
// ColorSilver
// ColorGray
// ColorRed
// ColorLime
// ColorYellow
// ColorBlue
// ColorFuchsia
// ColorAqua
// ColorWhite
