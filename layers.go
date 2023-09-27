package main

type Layers struct {
    rows   int
    cols   int
    ls     []*Layer
    active *Layer
}

func makeLayers(rows, cols int) *Layers {
    l := makeLayer(rows, cols)

    var ls = Layers{
        cols:   cols,
        rows:   rows,
        ls:     []*Layer{l},
        active: l,
    }

    return &ls
}

func (l Layers) setAt(c Cell) *Cell {
    return l.active.setAt(c)
}
