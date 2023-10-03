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

func (s HumanSector) GetText(x string, y int) *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "H", canvas.Center)
}

func (s HumanSector) SetX(x int) {
	s.x = x
}

func (s HumanSector) SetY(y int) {
	s.y = y
}
