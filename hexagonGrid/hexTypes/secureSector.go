package hexTypes

import (
	"fmt"
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

func (s SecureSector) GetText(x string, y int) *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, fmt.Sprintf("%s%02d", x, y+1), canvas.Center)
}

func (s SecureSector) SetX(x int) {
	s.x = x
}

func (s SecureSector) SetY(y int) {
	s.y = y
}
