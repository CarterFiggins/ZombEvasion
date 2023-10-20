package hexagonGrid

import (
	"math"
	"image/color"

	"infection/hexagonGrid/hexSectors"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

var (
	Board = &GameBoard{
		Loaded: false,
	}
)

type GameBoard struct {
	Grid [][]Hex
	Name string
	Loaded bool
	HumanSector *hexSectors.Human
	ZombieSector *hexSectors.Zombie
	CurrentDiscordUserID string
}

func (g *GameBoard) LoadBoard() error {
	g.Grid = MainBoard()
	g.Name = "main board"
	g.Loaded = true
	return CreateGameGrid(g.Grid)
}

func (g *GameBoard) UnloadGame() {
	Board = &GameBoard{
		Loaded: false,
	}
}

func CreateGameGrid(board [][]Hex) error {
	var canvasSizeX float64 = 100
	var canvasSizeY  float64 = 100

	c := canvas.New(canvasSizeX, canvasSizeY)
	ctx := canvas.NewContext(c)

	boardSizeY := len(board[0])
	boardSizeX := len(board)
	var height float64 = ((canvasSizeY - 4) / float64(boardSizeY)) / 2
	var hexRadius float64 = (2 * height) / math.Sqrt(3)
	var xHexRadius float64 = ((canvasSizeX - 2) /float64(boardSizeX)) / 1.57

	if xHexRadius < hexRadius {
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
		if xIndex % 2 == 0 {
			y = setY + height
		} else {
			y = setY
		}
		for yIndex := 0; yIndex <= boardSizeY - 1; yIndex++ {
			hex := board[xIndex][yIndex]
			hex.SetLocation(xIndex, yIndex)
			hexName := hex.GetSectorName()
			if (hexName == hexSectors.HumanSectorName) {
				Board.HumanSector = hex.(*hexSectors.Human)
			}
			if (hexName == hexSectors.ZombieName) {
				Board.ZombieSector = hex.(*hexSectors.Zombie)
			}
			drawHex(ctx, x, y, hexRadius, strokeWidth, hex.GetColor(), hex.GetStrokeColor())
			text, err := hex.GetText()
			if err != nil {
				return err
			}
			ctx.DrawText(x, y, text)
			y =  y - height * 2
		}
	}

	renderers.Write("gameBoard.png", c, canvas.DPMM(8.0))
	return nil
}


func drawHex(ctx *canvas.Context, x, y, radius, strokeWidth float64, color color.Color, lineColor color.Color) {
	ctx.SetStrokeColor(lineColor)
	ctx.SetStrokeWidth(strokeWidth)
	ctx.SetFillColor(color)
	ctx.DrawPath(x, y, canvas.RegularPolygon(6, radius, false))
}