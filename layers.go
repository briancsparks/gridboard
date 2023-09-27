package main

import (
    "github.com/gdamore/tcell/v2"
)

type Layers struct {
    rows   int
    cols   int
    ls     []*Layer
    active *Layer

    screen  tcell.Screen
}

func makeLayers(rows, cols int, sc tcell.Screen) *Layers {
    l := makeLayer(rows, cols)

    var ls = Layers{
        cols:   cols,
        rows:   rows,
        ls:     []*Layer{l},
        active: l,
        screen: sc,
    }

    return &ls
}

func (l Layers) setAt(cin Cell) *Cell {
    c := l.active.setAt(cin)
    if c == nil {
        return c
    }

    style := tcell.StyleDefault

    color, err := colorCode(c.Color)
    if err == nil {
       style = style.Foreground(color)
    }

    l.screen.SetCell(c.X, c.Y, style, c.Sym)

    if c.Redraw {
        l.screen.Show()
    }

    return c
}
