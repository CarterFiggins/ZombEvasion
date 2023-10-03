package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type EscapeSector struct {
	x int
	y int
	EscapeNumber int
}

func (s EscapeSector) GetColor() color.Color {
	return canvas.Black
}

func (s EscapeSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s EscapeSector) SetX(x int) {
	s.x = x
}

func (s EscapeSector) SetY(y int) {
	s.y = y
}
