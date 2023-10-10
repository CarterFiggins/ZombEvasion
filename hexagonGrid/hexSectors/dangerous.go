package hexSectors

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type Dangerous struct {
	Name string
	*Location
}

func (s Dangerous) GetColor() color.Color {
	return canvas.Gray
}

func (s Dangerous) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *Dangerous) GetSectorName() string {
	s.Name = DangerousName
	return DangerousName
}

func (s Dangerous) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, s.Location.GetHexName(), canvas.Center), nil
}

func (s *Dangerous) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}

func (s *Dangerous) CanMoveHere() bool {
	return true
}
