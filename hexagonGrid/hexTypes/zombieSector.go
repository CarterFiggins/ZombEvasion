package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type ZombieSector struct {
	col string
	row int
}

func (s ZombieSector) GetColor() color.Color {
	return canvas.Lawngreen 
}

func (s ZombieSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s ZombieSector) GetText() *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "Z", canvas.Center)
}

func (s *ZombieSector) SetCol(col string) {
	s.col = col
}

func (s *ZombieSector) SetRow(row int) {
	s.row = row
}
