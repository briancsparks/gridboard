package main

const (
    fX = 1 << iota
    fY
    fSym
    fColor
    fBgColor
)

const fAll uint32 = fX | fY | fSym | fColor | fBgColor

type Cell struct {
    X       int    `json:"x"`
    Y       int    `json:"y"`
    Sym     rune   `json:"sym"`
    Color   string `json:"color"`
    BgColor string `json:"bgcolor"`
    Redraw  bool   `json:"redraw"`

    flags   uint32
}

func makeCell(x, y int, sym rune) *Cell {
    var c = Cell{
        X:       x,
        Y:       y,
        Sym:     sym,
        Redraw:  false,
        flags:   fX | fY | fSym,
    }

    return &c
}

func makeEmptyCell(x, y int) *Cell {
    var c = Cell{
        X:       x,
        Y:       y,
        Redraw:  false,
        flags:   fX | fY,
    }

    return &c
}

func makeFullCell(x, y int, sym rune, color, bgcolor string) *Cell {
    var c = Cell{
        X:       x,
        Y:       y,
        Sym:     sym,
        Color:   color,
        BgColor: bgcolor,
        Redraw:  false,
        flags:   fX | fY | fSym | fColor | fBgColor,
    }

    return &c
}

func (c *Cell) accumulate(that *Cell) bool {
    if !c.isComplete() {
        if !c.hasX() && that.hasX() {
            c.setX(that.X)
        }
        if !c.hasY() && that.hasY() {
            c.setY(that.Y)
        }
        if !c.hasSym() && that.hasSym() {
            c.setSym(that.Sym)
        }
        if !c.hasColor() && that.hasColor() {
            c.setColor(that.Color)
        }
        if !c.hasBgColor() && that.hasBgColor() {
            c.setBgColor(that.BgColor)
        }
    }

    return c.isComplete()
}

func (c *Cell) isComplete() bool {
    return c.flags == fAll
}

func (c *Cell) isEmpty() bool {
    return !c.hasSym() && !c.hasColor() && !c.hasBgColor()
}

func (c *Cell) hasX() bool {
    return (c.flags & fX) != 0
}

func (c *Cell) hasY() bool {
    return (c.flags & fY) != 0
}

func (c *Cell) hasSym() bool {
    return (c.flags & fSym) != 0
}

func (c *Cell) hasColor() bool {
    return (c.flags & fColor) != 0
}

func (c *Cell) hasBgColor() bool {
    return (c.flags & fBgColor) != 0
}


func (c *Cell) setX(x int) *Cell {
    c.X = x
    c.flags |= fX
    return c
}

func (c *Cell) setY(y int) *Cell {
    c.Y = y
    c.flags |= fY
    return c
}

func (c *Cell) setSym(r rune) *Cell {
    c.Sym = r
    c.flags |= fSym
    return c
}

func (c *Cell) setColor(s string) *Cell {
    c.Color = s
    c.flags |= fColor
    return c
}

func (c *Cell) setBgColor(s string) *Cell {
    c.BgColor = s
    c.flags |= fBgColor
    return c
}

func (c *Cell) x() *int {
    if (c.flags & fX) != 0 {
        return &c.X
    }
    return nil
}

func (c *Cell) y() *int {
    if (c.flags & fY) != 0 {
        return &c.Y
    }
    return nil
}

func (c *Cell) sym() *rune {
    if (c.flags & fSym) != 0 {
        return &c.Sym
    }
    return nil
}

func (c *Cell) color() string {
    if (c.flags & fColor) != 0 {
        return c.Color
    }
    return ""
}

func (c *Cell) bgcolor() string {
    if (c.flags & fBgColor) != 0 {
        return c.BgColor
    }
    return ""
}

