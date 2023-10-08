package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type HumanSector struct {
	name string
	col int
	row int
}

func (s HumanSector) GetColor() color.Color {
	return canvas.Navy 
}

func (s HumanSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *HumanSector) GetSectorName() string {
	s.name = HumanSectorName
	return HumanSectorName
}

func (s HumanSector) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(8.0, canvas.White, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "H", canvas.Center), nil
}

func (s *HumanSector) SetCol(col int) {
	s.col = col
}

func (s *HumanSector) SetRow(row int) {
	s.row = row
}
