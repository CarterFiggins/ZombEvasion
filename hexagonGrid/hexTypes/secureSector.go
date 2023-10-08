package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type SecureSector struct {
	Name string
	Col int
	Row int
}

func (s SecureSector) GetColor() color.Color {
	return canvas.White
}

func (s SecureSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *SecureSector) GetSectorName() string {
	s.Name = SecureSectorName
	return SecureSectorName
}

func (s SecureSector) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	hexName, err := HexName(s.Col, s.Row+1)
	if err != nil {
		return nil, err
	}

	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, hexName, canvas.Center), nil
}

func (s *SecureSector) SetCol(col int) {
	s.Col = col
}

func (s *SecureSector) SetRow(row int) {
	s.Row = row
}