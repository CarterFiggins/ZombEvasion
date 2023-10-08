package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type WallSector struct {
	Name string
	Col int
	Row int
}

func (s WallSector) GetColor() color.Color {
	return canvas.Transparent
}

func (s WallSector) GetStrokeColor() color.Color {
	return canvas.Transparent
}

func (s *WallSector) GetSectorName() string {
	s.Name = WallSectorName
	return WallSectorName
}

func (s WallSector) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "", canvas.Center), nil
}

func (s *WallSector) SetCol(col int) {
	s.Col = col
}

func (s *WallSector) SetRow(row int) {
	s.Row = row
}
