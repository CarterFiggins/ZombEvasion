package hexSectors

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type Secure struct {
	Name string
	*Location
}

func (s Secure) GetColor() color.Color {
	return canvas.White
}

func (s Secure) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *Secure) GetSectorName() string {
	s.Name = SecureName
	return SecureName
}

func (s Secure) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, s.Location.GetHexName(), canvas.Center), nil
}

func (s *Secure) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}

func (s *Secure) CanMoveHere() bool {
	return true
}
