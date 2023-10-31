package hexagonGrid

import (
	"image/color"

	"github.com/tdewolff/canvas"
	"infection/hexagonGrid/hexSectors"
)

type Hex interface {
	GetColor() color.Color
	GetStrokeColor() color.Color
	GetText() (*canvas.Text, error)
	SetLocation(int, int)
	GetSectorName() string
	CanMoveHere() bool
}

func TestBoard() [][]Hex{
	return [][]Hex{
		{
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
		},
		{
			&hexSectors.SafeHouse{EscapeNumber: 1},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Zombie{},
		},
		{
			&hexSectors.Dangerous{},
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
			&hexSectors.Human{},
		},
	}
}
