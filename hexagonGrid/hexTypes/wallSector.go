package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type WallSector struct {
	col string
	row int
}

func (s WallSector) GetColor() color.Color {
	return canvas.Transparent
}

func (s WallSector) GetStrokeColor() color.Color {
	return canvas.Transparent
}

func (s WallSector) GetText() *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "", canvas.Center)
}

func (s *WallSector) SetCol(col string) {
	s.col = col
}

func (s *WallSector) SetRow(row int) {
	s.row = row
}
