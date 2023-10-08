package hexagonGrid

import (
	"image/color"

	"github.com/tdewolff/canvas"
	"infection/hexagonGrid/hexTypes"
)

type Hex interface {
	GetColor() color.Color
	GetStrokeColor() color.Color
	GetText() *canvas.Text
	SetCol(string)
	SetRow(int)
	GetSector() interface{}
}

func MainBoard() [][]Hex{
	return [][]Hex{
		// a
		{
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
		},
		// b
		{
			&hexTypes.DangerousSector{},
			&hexTypes.EscapeSector{EscapeNumber: 1},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.EscapeSector{EscapeNumber: 4},
			&hexTypes.DangerousSector{},
		},
		// c
		{
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// d
		{
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// e
		{
			&hexTypes.WallSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
		},
		// f
		{
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
		},
		// g
		{
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// h
		{
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// i
		{
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// j
		{
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// k
		{
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
		},
		// l
		{
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.ZombieSector{},
			&hexTypes.WallSector{},
			&hexTypes.HumanSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// m
		{
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
		},
		// n
		{
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// o
		{
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// p
		{
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
		},
		// q
		{
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
		},
		// r
		{
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
		},
		// s
		{
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
		},
		// t
		{
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
		},
		// u
		{
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
		},
		// v
		{
			&hexTypes.SecureSector{},
			&hexTypes.EscapeSector{EscapeNumber: 2},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.WallSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.EscapeSector{EscapeNumber: 3},
			&hexTypes.DangerousSector{},
		},
		// w
		{
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.DangerousSector{},
		},
	}
}

func TestBoard() [][]Hex{
	return [][]Hex{
		{
			&hexTypes.WallSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.EscapeSector{EscapeNumber: 1},
		},
		{
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.WallSector{},
			&hexTypes.ZombieSector{},
		},
		{
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
			&hexTypes.SecureSector{},
			&hexTypes.HumanSector{},
		},
		{
			&hexTypes.SecureSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
		},
		{
			&hexTypes.SecureSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
		},
		{
			&hexTypes.SecureSector{},
			&hexTypes.WallSector{},
			&hexTypes.WallSector{},
			&hexTypes.DangerousSector{},
		},
		{
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
		},
		{
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
		},
		{
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
		},
		{
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.SecureSector{},
			&hexTypes.DangerousSector{},
		},
		
	}
}