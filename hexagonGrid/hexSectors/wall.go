package hexSectors

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type Wall struct {
	Name string
	*Location
}

func (s Wall) GetColor() color.Color {
	return canvas.Transparent
}

func (s Wall) GetStrokeColor() color.Color {
	return canvas.Transparent
}

func (s *Wall) GetSectorName() string {
	s.Name = WallName
	return WallName
}

func (s Wall) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "", canvas.Center), nil
}

func (s *Wall) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}

func (s *Wall) CanMoveHere() bool {
	return false
}
