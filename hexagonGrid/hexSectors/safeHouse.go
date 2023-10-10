package hexSectors

import (
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
)

type SafeHouse struct {
	Name string
	*Location
	EscapeNumber int
}

func (s SafeHouse) GetColor() color.Color {
	return canvas.Green
}

func (s SafeHouse) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *SafeHouse) GetSectorName() string {
	s.Name = SafeHouseName
	return SafeHouseName
}

func (s SafeHouse) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(4.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, fmt.Sprintf("%s\n%d", s.Location.GetHexName(), s.EscapeNumber), canvas.Center), nil
}

func (s *SafeHouse) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}

func (s *SafeHouse) CanMoveHere() bool {
	return true
}
