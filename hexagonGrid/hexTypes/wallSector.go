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

func (s WallSector) GetText(x string, y int) *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "", canvas.Center)
}

func (s WallSector) SetX(x int) {
	s.x = x
}

func (s WallSector) SetY(y int) {
	s.y = y
}
