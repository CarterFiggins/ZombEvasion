package hexTypes

import (
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
)

type DangerousSector struct {
	x int
	y int
}

func (s DangerousSector) GetColor() color.Color {
	return canvas.Gray
}

func (s DangerousSector) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s DangerousSector) GetText(x string, y int) *canvas.Text {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		panic(err)
	}
	face := fontFamily.Face(5.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, fmt.Sprintf("%s%02d", x, y+1), canvas.Center)
}
