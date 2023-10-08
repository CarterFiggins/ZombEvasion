package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type ZombieSector struct {
	name string
	col int
	row int
}

func (s ZombieSector) GetColor() color.Color {
	return canvas.Lawngreen 
}

func (s ZombieSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *ZombieSector) GetSectorName() string {
	s.name = ZombieSectorName
	return ZombieSectorName
}

func (s ZombieSector) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "Z", canvas.Center), nil
}

func (s *ZombieSector) SetCol(col int) {
	s.col = col
}

func (s *ZombieSector) SetRow(row int) {
	s.row = row
}
