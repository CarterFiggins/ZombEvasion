package hexTypes

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type HumanSector struct {
	col string
	row int
}

func (s HumanSector) GetColor() color.Color {
	return canvas.Navy 
}

func (s HumanSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s HumanSector) GetText() *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(8.0, canvas.White, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "H", canvas.Center)
}

func (s *HumanSector) SetCol(col string) {
	s.col = col
}

func (s *HumanSector) SetRow(row int) {
	s.row = row
}
