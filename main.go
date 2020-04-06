package main

import (
    board "github.com/kyrozetera/eight-queens-go/queensboard"
)

func main() {
    mainboard := board.NewBoard()
    mainboard.Generate()
    mainboard.Print()
}
