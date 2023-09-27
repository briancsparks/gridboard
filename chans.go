package main

type Chans struct {
    cell chan Cell
    quit chan bool
}

func makeChans() *Chans {
    var c Chans

    c.quit = make(chan bool)
    c.cell = make(chan Cell, 100)

    return &c
}
