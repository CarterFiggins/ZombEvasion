package hexTypes

import (
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
)

type EscapeSector struct {
	name string
	col int
	row int
	EscapeNumber int
}

func (s EscapeSector) GetColor() color.Color {
	return canvas.Green
}

func (s EscapeSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *EscapeSector) GetSectorName() string {
	s.name = EscapeSectorName
	return EscapeSectorName
}

func (s EscapeSector) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(6.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, fmt.Sprintf("%d", s.EscapeNumber), canvas.Center), nil
}

func (s *EscapeSector) SetCol(col int) {
	s.col = col
}

func (s *EscapeSector) SetRow(row int) {
	s.row = row
}
