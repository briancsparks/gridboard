package grid

type Chans struct {
    Cell chan Cell
    Quit chan bool
}

func MakeChans() *Chans {
    var c Chans

    c.Quit = make(chan bool)
    c.Cell = make(chan Cell, 100)

    return &c
}
