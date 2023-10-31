package hexagonGrid

import (
	"infection/hexagonGrid/hexSectors"
)

func SmallBoard() [][]Hex{
	return [][]Hex{
		{
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
		},
		{
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Dangerous{},
			&hexSectors.Wall{},
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Dangerous{},
			&hexSectors.Wall{},
		},
		{
			&hexSectors.SafeHouse{EscapeNumber: 3},
			&hexSectors.Wall{},
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
			&hexSectors.SafeHouse{EscapeNumber: 1},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
		},
		{
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Dangerous{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Dangerous{},
		},
		{
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Dangerous{},
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Zombie{},
			&hexSectors.Human{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
		},
		{
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
			&hexSectors.Dangerous{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
		},
		{
			&hexSectors.SafeHouse{EscapeNumber: 4},
			&hexSectors.Wall{},
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
			&hexSectors.SafeHouse{EscapeNumber: 2},
			&hexSectors.Wall{},
			&hexSectors.Dangerous{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Dangerous{},
		},
		{
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Dangerous{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Wall{},
		},
		{
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Wall{},
			&hexSectors.Secure{},
			&hexSectors.Dangerous{},
			&hexSectors.Wall{},
		},
	}
}