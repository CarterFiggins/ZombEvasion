package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type SecureSector struct {
	name string
	col int
	row int
}

func (s SecureSector) GetColor() color.Color {
	return canvas.White
}

func (s SecureSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *SecureSector) GetSectorName() string {
	s.name = SecureSectorName
	return SecureSectorName
}

func (s SecureSector) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	hexName, err := HexName(s.col, s.row+1)
	if err != nil {
		return nil, err
	}

	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, hexName, canvas.Center), nil
}

func (s *SecureSector) SetCol(col int) {
	s.col = col
}

func (s *SecureSector) SetRow(row int) {
	s.row = row
}