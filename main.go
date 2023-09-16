package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "github.com/gdamore/tcell/v2"
    "os"
)

type Cell struct {
	X   int  `json:"x"`
	Y   int  `json:"y"`
	Sym rune `json:"sym"`
	Color  string `json:"color"`
	Redraw bool   `json:"redraw"`
}

func readFifo(line chan string) {
	for {
    f, err := os.OpenFile("/tmp/termgridboard", os.O_RDONLY, 0600)
    if err != nil {
      panic(err)
    }

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
      line <- scanner.Text()
    }
    f.Close()
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

	//scanner := bufio.NewScanner(f)

	quit := make(chan bool)
	line := make(chan string, 100)

    go readFifo(line)

	//// Goroutine to read from the FIFO
	//go func() {
	//	for scanner.Scan() {
	//		line <- scanner.Text()
	//	}
	//	//fmt.Println("Scanner done.")
	//}()

    problem_json := ""

	go func() {
		for {
			select {
			case <-quit:
				return

			case l := <-line:
				var cell Cell
				if err := json.Unmarshal([]byte(l), &cell); err != nil {
                    problem_json = l
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
                fmt.Println(problem_json)
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
