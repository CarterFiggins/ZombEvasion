package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type DangerousSector struct {
	x int
	y int
}

func (s DangerousSector) GetColor() color.Color {
	return canvas.Gray
}

func (s DangerousSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s DangerousSector) SetX(x int) {
	s.x = x
}

func (s DangerousSector) SetY(y int) {
	s.y = y
}
