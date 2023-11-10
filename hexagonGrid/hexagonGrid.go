package hexagonGrid

import (
	"math"
	"image/color"
	"fmt"
	"errors"

	"infection/models"
	"infection/hexagonGrid/hexSectors"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

const (
	ForestBoardName string = "forestBoard"
	GraveYardBoardName = "graveYardBoard"
	HospitalBoardName = "hospitalBoard"
)

func GetBoard(guildID string) ([][]Hex, error) {
	game, err := models.GetGame(guildID)
	if err != nil {
		return [][]Hex{}, err
	}
	if game.BoardName == ForestBoardName {
		return ForestBoard(), nil
	}
	if game.BoardName == GraveYardBoardName {
		return GraveYardBoard(), nil
	}
	if game.BoardName == HospitalBoardName {
		return HospitalBoard(), nil
	}
	return [][]Hex{}, errors.New("No board found! Might need to run `/server-setup`")
}

func FindZombieSector(board [][]Hex) *hexSectors.Zombie {
	for xIndex := 0; xIndex < len(board); xIndex++ {
		for yIndex := 0; yIndex < len(board[0]); yIndex++ {
			hex := board[xIndex][yIndex]
			hex.SetLocation(xIndex, yIndex)
			hexName := hex.GetSectorName()
			if (hexName == hexSectors.ZombieName) {
				return hex.(*hexSectors.Zombie)
			}
		}
	}
	return nil
}

func FindHumanSector(board [][]Hex) *hexSectors.Human {
	for xIndex := 0; xIndex < len(board); xIndex++ {
		for yIndex := 0; yIndex < len(board[0]); yIndex++ {
			hex := board[xIndex][yIndex]
			hex.SetLocation(xIndex, yIndex)
			hexName := hex.GetSectorName()
			if (hexName == hexSectors.HumanName) {
				return hex.(*hexSectors.Human)
			}
		}
	}
	return nil
}

func CreateAllGameImages() error {
	err := CreateGameGridImage(ForestBoard(), ForestBoardName)
	if err != nil {
		return err
	}
	err = CreateGameGridImage(GraveYardBoard(), GraveYardBoardName)
	if err != nil {
		return err
	}
	err = CreateGameGridImage(HospitalBoard(), HospitalBoardName)
	if err != nil {
		return err
	}
	return nil
}

func CreateGameGridImage(board [][]Hex, fileName string) error {
	var canvasSizeX float64 = 100
	var canvasSizeY  float64 = 100

	goCanvas := canvas.New(canvasSizeX, canvasSizeY)
	ctx := canvas.NewContext(goCanvas)

	boardSizeY := len(board[0])
	boardSizeX := len(board)
	var height float64 = ((canvasSizeY) / (float64(boardSizeY)+.5)) / 2
	var hexRadius float64 = (2 * height) / math.Sqrt(3)
	var xHexRadius float64 = ((canvasSizeX) /float64(boardSizeX)) / 1.57
	setY := canvasSizeY - (height)*2

	if xHexRadius < hexRadius {
		height = (xHexRadius * math.Sqrt(3))/2
		setY = canvasSizeY - (height) * 2
		hexRadius = xHexRadius
	}
	var strokeWidth float64 = hexRadius / 10
	var setX float64 = -hexRadius/3
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
			drawHex(ctx, x, y, hexRadius, strokeWidth, hex.GetColor(), hex.GetStrokeColor())
			text, err := hex.GetText()
			if err != nil {
				return err
			}
			ctx.DrawText(x, y, text)
			y =  y - height * 2
		}
	}

	renderers.Write(fmt.Sprintf("%s.png", fileName), goCanvas, canvas.DPMM(8.0))
	return nil
}


func drawHex(ctx *canvas.Context, x, y, radius, strokeWidth float64, color color.Color, lineColor color.Color) {
	ctx.SetStrokeColor(lineColor)
	ctx.SetStrokeWidth(strokeWidth)
	ctx.SetFillColor(color)
	ctx.DrawPath(x, y, canvas.RegularPolygon(6, radius, false))
}