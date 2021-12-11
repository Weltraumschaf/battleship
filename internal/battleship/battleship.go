package battleship

import (
	"bytes"
	"fmt"
)

type Coordinate struct {
	row    ColumnCoordinate
	column RowCoordinate
}

type ColumnCoordinate string

const (
	A ColumnCoordinate = "A"
	B                  = "B"
	C = "C"
	D = "D"
	E = "E"
	F = "F"
	G = "G"
	H = "H"
	I = "I"
	J = "J"
)

type RowCoordinate int

const (
	One RowCoordinate = iota + 1
	Two 
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
)

type Board struct {
	
}

func NewBoard() *Board {
	return &Board{}
}

func (b Board) String() string {
	var buf = bytes.Buffer{}

	buf.WriteString("   | A | B | C | D | E | F | G | H |I  | J |\n")
	buf.WriteString("---+---+---+---+---+---+---+---+---+---+---+\n")

	for i := 1; i <= 10; i++ {
		formatRow(&buf, i)
	}

	return buf.String()
}

func formatRow(buf *bytes.Buffer, rowNumber int) {
	const rowFormat = "%2d |   |   |   |   |   |   |   |   |   |   |\n"
	buf.WriteString(fmt.Sprintf(rowFormat, rowNumber))
	buf.WriteString("---+---+---+---+---+---+---+---+---+---+---+\n")
}
