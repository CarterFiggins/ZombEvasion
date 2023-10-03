package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type HumanSector struct {
	x int
	y int
}

func (s HumanSector) GetColor() color.Color {
	return canvas.Blue
}

func (s HumanSector) GetStrokeColor() color.Color {
	return canvas.Black
}


func (s HumanSector) SetX(x int) {
	s.x = x
}

func (s HumanSector) SetY(y int) {
	s.y = y
}
