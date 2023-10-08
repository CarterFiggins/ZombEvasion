package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type DangerousSector struct {
	name string
	col int
	row int
}

func (s DangerousSector) GetColor() color.Color {
	return canvas.Gray
}

func (s DangerousSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *DangerousSector) GetSectorName() string {
	s.name = DangerousSectorName
	return DangerousSectorName
}

func (s DangerousSector) GetText() (*canvas.Text, error) {
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

func (s *DangerousSector) SetCol(col int) {
	s.col = col
}

func (s *DangerousSector) SetRow(row int) {
	s.row = row
}