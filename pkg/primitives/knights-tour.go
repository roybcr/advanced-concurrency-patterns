package primitives

import (
	"fmt"
)

const BOARD_SIZE int = 8

type Square struct {
	visited bool
	x       int
	y       int
}

type Board  struct { board [BOARD_SIZE][BOARD_SIZE]Square }
type Knight struct { sqr Square }

type Moveable interface {
	up(brd *Board) []Square
	down(brd *Board) []Square
	left(brd *Board) []Square
	right(brd *Board) []Square
}

func NewSquare(x, y int) Square  { return Square{x: x, y: y} }

func (s *Square) ToggleVisited() { (*s).visited = !(*s).visited }

func (s *Square) OnEdge() bool {
	sqr := *s
	return (sqr.x == 0 || sqr.x == 7 || sqr.y == 0 || sqr.y == 7)
}

func (s *Square) OnCorner() bool {
	if !(*s).OnEdge() {
		return false
	}

	sqr := *s
	return (sqr.x == sqr.y || sqr.x == 0 && sqr.y == 7 || sqr.x == 7 && sqr.y == 0)
}

func (b *Board) New() Board {

	for y, row := range (*b).board {
		for x, col := range row {
			col = NewSquare(x, y)
			row[x] = col
		}

		(*b).board[y] = row
	}

	return *b
}

func (b *Board) Print() {
	for _, row := range (*b).board {
		for _, col := range row {
			fmt.Printf(" [ %v ] ", col)
		}

		fmt.Println()
	}
}

func (b *Board) FindSquares(kn *Knight) []Square {

	sqr := (*kn).sqr
	x := sqr.x
	y := sqr.y

	squares := []Square{}

	if sqr.OnCorner() {

		if y == 7 {

			if x == 7 {
				sqr1 := (*b).board[x-2][y-1]
				sqr2 := (*b).board[x-1][y-2]
				squares = append([]Square{}, sqr1, sqr2)
			} else {
				sqr1 := (*b).board[x+2][y-1]
				sqr2 := (*b).board[x+1][y-2]
				squares = append([]Square{}, sqr1, sqr2)
			}

		} else {

			if x == 7 {
				sqr1 := (*b).board[x-2][y+1]
				sqr2 := (*b).board[x-1][y+2]
				squares = append([]Square{}, sqr1, sqr2)
			} else {
				sqr1 := (*b).board[x+2][y+1]
				sqr2 := (*b).board[x+1][y+2]
				squares = append([]Square{}, sqr1, sqr2)
			}
		}

	}

	availableSquares := []Square{}

	for i := range squares {
		if !squares[i].visited {
			availableSquares = append([]Square(availableSquares), squares[i])
		}
	}

	return availableSquares

}

func KnightsTour() {

	var board Board

	board = board.New()
	board.Print()

	for _, row := range board.board {
		for _, col := range row {
			kn := Knight{sqr: col}

			if kn.sqr.OnCorner() {
				fmt.Printf("For a Knight on Square [%d, %d], Possible Available Squares are: %v\n", kn.sqr.x, kn.sqr.y, board.FindSquares(&kn))
			}
		}
	}
}
