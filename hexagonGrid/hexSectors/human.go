package hexSectors

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type Human struct {
	Name string
	*Location
}

func (s Human) GetColor() color.Color {
	return canvas.Navy 
}

func (s Human) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *Human) GetSectorName() string {
	s.Name = HumanSectorName
	return HumanSectorName
}

func (s Human) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(8.0, canvas.White, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "H", canvas.Center), nil
}

func (s *Human) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}

func (s *Human) CanMoveHere() bool {
	return false
}
