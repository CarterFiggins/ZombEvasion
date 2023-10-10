package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type DangerousSector struct {
	Name string
	*Location
}

func (s DangerousSector) GetColor() color.Color {
	return canvas.Gray
}

func (s DangerousSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *DangerousSector) GetSectorName() string {
	s.Name = DangerousSectorName
	return DangerousSectorName
}

func (s DangerousSector) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, s.Location.GetHexName(), canvas.Center), nil
}

func (s *DangerousSector) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}

func (s *DangerousSector) CanMoveHere() bool {
	return true
}
