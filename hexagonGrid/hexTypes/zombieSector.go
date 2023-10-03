package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type ZombieSector struct {
	x int
	y int
}

func (s ZombieSector) GetColor() color.Color {
	return canvas.Purple
}

func (s ZombieSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s ZombieSector) GetText(x string, y int) *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "Z", canvas.Center)
}

func (s ZombieSector) SetX(x int) {
	s.x = x
}

func (s ZombieSector) SetY(y int) {
	s.y = y
}
