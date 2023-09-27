package main

type Chans struct {
    cell chan Cell
}

func makeChans() *Chans {
    var c Chans

    c.cell = make(chan Cell, 100)

    return &c
}
