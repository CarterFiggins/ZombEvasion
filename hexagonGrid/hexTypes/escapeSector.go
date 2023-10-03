package hexTypes

import (
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
)

type EscapeSector struct {
	x int
	y int
	EscapeNumber int
}

func (s EscapeSector) GetColor() color.Color {
	return canvas.Green
}

func (s EscapeSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s EscapeSector) GetText(x string, y int) *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(6.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, fmt.Sprintf("E%d", s.EscapeNumber), canvas.Center)
}

func (s EscapeSector) SetX(x int) {
	s.x = x
}

func (s EscapeSector) SetY(y int) {
	s.y = y
}
