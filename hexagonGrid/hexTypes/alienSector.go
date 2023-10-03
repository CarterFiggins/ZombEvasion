package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type AlienSector struct {
	x int
	y int
}

func (s AlienSector) GetColor() color.Color {
	return canvas.Purple
}

func (s AlienSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s AlienSector) SetX(x int) {
	s.x = x
}

func (s AlienSector) SetY(y int) {
	s.y = y
}
