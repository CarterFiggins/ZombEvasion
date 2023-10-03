package hexTypes

import (
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
)

type SecureSector struct {
	col string
	row int
}

func (s SecureSector) GetColor() color.Color {
	return canvas.White
}

func (s SecureSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s SecureSector) GetText() *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, fmt.Sprintf("%s%02d", s.col, s.row+1), canvas.Center)
}

func (s *SecureSector) SetCol(col string) {
	s.col = col
}

func (s *SecureSector) SetRow(row int) {
	s.row = row
}