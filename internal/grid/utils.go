package grid

import (
    "fmt"
    "github.com/gdamore/tcell/v2"
)

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
