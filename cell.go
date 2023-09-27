package main

type Cell struct {
X      int    `json:"x"`
Y      int    `json:"y"`
Sym    rune   `json:"sym"`
Color  string `json:"color"`
Redraw bool   `json:"redraw"`
}
