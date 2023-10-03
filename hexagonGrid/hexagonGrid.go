package hexagonGrid

import (
	"math"
	"image/color"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

func CreateGrid(board [][]Hex) {
	var canvasSizeX float64 = 100
	var canvasSizeY  float64 = 100
	letterMap := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
		5: "F",
		6: "G",
		7: "H",
		8: "I",
		9: "J",
		10: "K",
		11: "L",
		12: "M",
		13: "N",
		14: "O",
		15: "P",
		16: "Q",
		17: "R",
		18: "S",
		19: "T",
		20: "U",
		21: "V",
		22: "W",
		23: "X",
		24: "Y",
		25: "Z",
	}
	c := canvas.New(canvasSizeX, canvasSizeY)
	ctx := canvas.NewContext(c)

	boardSizeY := len(board[0])
	boardSizeX := len(board)
	var height float64 = ((canvasSizeY - 4) / float64(boardSizeY)) / 2
	var hexRadius float64 = (2 * height) / math.Sqrt(3)
	var xHexRadius float64 = ((canvasSizeX - 2) /float64(boardSizeX)) / 1.57

	if (xHexRadius < hexRadius) {
		height = (xHexRadius * math.Sqrt(3))/2
		hexRadius = xHexRadius
	}
	var strokeWidth float64 = hexRadius / 10
	setY := canvasSizeY - height*2
	var setX float64 = 0
	x := setX
	y := setY

	for xIndex := 0; xIndex <= boardSizeX - 1; xIndex++ {
		x = x + hexRadius * 1.50
		if (xIndex % 2 == 0) {
			y = setY + height
			} else {
			y = setY
		}
		for yIndex := 0; yIndex <= boardSizeY - 1; yIndex++ {
			hex := board[xIndex][yIndex]
			drawHex(ctx, x, y, hexRadius, strokeWidth, hex.GetColor(), hex.GetStrokeColor())
			letter, ok := letterMap[xIndex]
			if !ok {
				panic("no letter found in map")
			}
			ctx.DrawText(x, y, hex.GetText(letter, yIndex))
			y =  y - height * 2
		}
	}

	renderers.Write("out.png", c, canvas.DPMM(8.0))
}


func drawHex(ctx *canvas.Context, x, y, radius, strokeWidth float64, color color.Color, lineColor color.Color) {
	ctx.SetStrokeColor(lineColor)
	ctx.SetStrokeWidth(strokeWidth)
	ctx.SetFillColor(color)
	ctx.DrawPath(x, y, canvas.RegularPolygon(6, radius, false))
}