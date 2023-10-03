package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type WallSector struct {
	x int
	y int
}

func (s WallSector) GetColor() color.Color {
	return canvas.Transparent
}

func (s WallSector) GetStrokeColor() color.Color {
	return canvas.Transparent
}

func (s WallSector) SetX(x int) {
	s.x = x
}

func (s WallSector) SetY(y int) {
	s.y = y
}
