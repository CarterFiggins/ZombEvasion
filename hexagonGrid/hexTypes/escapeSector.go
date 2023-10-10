package hexTypes

import (
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
)

type EscapeSector struct {
	Name string
	*Location
	EscapeNumber int
}

func (s EscapeSector) GetColor() color.Color {
	return canvas.Green
}

func (s EscapeSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *EscapeSector) GetSectorName() string {
	s.Name = EscapeSectorName
	return EscapeSectorName
}

func (s EscapeSector) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(4.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, fmt.Sprintf("%s\n%d", s.Location.GetHexName(), s.EscapeNumber), canvas.Center), nil
}

func (s *EscapeSector) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}

func (s *EscapeSector) CanMoveHere() bool {
	return true
}
