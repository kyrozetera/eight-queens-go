package queensboard

import (
    "fmt"
)

type Coordinate struct {
    X, Y int
}

type QueensBoard struct {
    Coords [][]int
    Queens []Coordinate
    Size Coordinate
}

func NewBoard() QueensBoard {
    board := QueensBoard {
        Coords: make([][]int, 8),
        Queens: make([]Coordinate, 0),
        Size: Coordinate{X: 8, Y: 8},
    }

    for y := range board.Coords {
        board.Coords[y] = make([]int, 8)
    }

    return board
}

func (board *QueensBoard) Generate() {
    success := board.addQueens()
    if (success) {
        fmt.Println("Success")
    } else {
        fmt.Println("Failed")
    }
}

func (board QueensBoard) Print() {
    for _, queen := range board.Queens {
        fmt.Printf("X: %d, Y: %d\n", queen.X, queen.Y)
    }

    for _, row := range board.Coords {
        for _, state := range row {
            if (state == -1) {
                fmt.Printf("Q ")
            }

            if (state > 0) {
                fmt.Printf("X ")
            }
        }
        fmt.Printf("\n")
    }
}

func (board *QueensBoard) addQueens() bool {
    if (len(board.Queens) == 8) {
        return true
    }

    for _, space := range board.available() {
        board.addQueen(space)

        success := board.addQueens()
        if success {
            return true
        }
        board.removeLastQueen()
    }

    return false
}

func (board *QueensBoard) addQueen(space Coordinate) {
    board.Queens = append(board.Queens, space)

    board.Coords[space.Y][space.X] = -1

    board.modifyAttacks(space, 1)
}

func (board *QueensBoard) modifyAttacks(space Coordinate, incrementBy int) {
    for x:=space.X+1; x < board.Size.X; x++ {
        board.Coords[space.Y][x] += incrementBy
    }

    for x:=space.X-1; x >= 0; x-- {
        board.Coords[space.Y][x] += incrementBy
    }

    for y:=space.Y+1; y < board.Size.Y; y++ {
        board.Coords[y][space.X] += incrementBy
    }

    for y:=space.Y-1; y >= 0; y-- {
        board.Coords[y][space.X] += incrementBy
    }

    y := space.Y+1
    for x:=space.X+1; x < board.Size.X; x++ {
        if (y >= board.Size.Y) {
            break
        }
        board.Coords[y][x] += incrementBy
        y++
    }

    y = space.Y-1
    for x:=space.X+1; x < board.Size.X; x++ {
        if (y < 0) {
            break
        }
        board.Coords[y][x] += incrementBy
        y--
    }

    y = space.Y+1
    for x:=space.X-1; x >= 0; x-- {
        if (y >= board.Size.Y) {
            break
        }
        board.Coords[y][x] += incrementBy
        y++
    }

    y = space.Y-1
    for x:=space.X-1; x >= 0; x-- {
        if (y < 0) {
            break
        }
        board.Coords[y][x] += incrementBy
        y--
    }
}

func (board *QueensBoard) removeLastQueen() {
    queen := board.Queens[len(board.Queens)-1]

    board.Queens = board.Queens[:len(board.Queens)-1]

    board.Coords[queen.Y][queen.X] = 0

    board.modifyAttacks(queen, -1)
}

func (board QueensBoard) available() []Coordinate {
    availableCoords := make([]Coordinate, 0)

    for y, row := range board.Coords {
        for x, state := range row {
            if state == 0 {
                availableCoords = append(availableCoords, Coordinate{X: x, Y: y})
            }
        }
    }

    return availableCoords
}
