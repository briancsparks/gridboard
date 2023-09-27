package main

type Layer struct {
    cells [][]Cell
}

func makeLayer(rows, cols int) *Layer {
    var l = Layer{
        cells: make([][]Cell, rows),
    }

    for i := 0; i < rows; i++ {
        l.cells[i] = make([]Cell, cols)
        for j := 0; j < cols; j++ {
            l.cells[i][j] = *makeEmptyCell(i, j)
        }
    }

    return &l
}

func (l Layer) at(x, y int) *Cell {
    if x >= len(l.cells) {
        return nil
    }
    if y >= len(l.cells[0]) {
        return nil
    }

    return &l.cells[x][y]
}

func (l Layer) setAt(c Cell) *Cell {
    if c.X >= len(l.cells) {
        return nil
    }
    if c.Y >= len(l.cells[0]) {
        return nil
    }

    l.cells[c.X][c.Y] = c
    return l.at(c.X, c.Y)
}
