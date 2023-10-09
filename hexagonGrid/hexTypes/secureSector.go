package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type SecureSector struct {
	Name string
	*Location
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

	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, s.Location.HexName(), canvas.Center), nil
}

func (s *SecureSector) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}

func (s *SecureSector) CanMoveHere() bool {
	return true
}
