package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type SecureSector struct {
	x int
	y int
}

func (s SecureSector) GetColor() color.Color {
	return canvas.White
}

func (s SecureSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s SecureSector) SetX(x int) {
	s.x = x
}

func (s SecureSector) SetY(y int) {
	s.y = y
}
