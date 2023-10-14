package game

import (
	"infection/hexagonGrid/hexSectors"
	"infection/hexagonGrid"
)

func CanSetOffAlarmHere(setOffX, setOffY int) bool {
	grid := hexagonGrid.Board.Grid
	if setOffX >= len(grid) || setOffX < 0 {
		return false
	} else if setOffY >= len(grid[setOffX]) || setOffY < 0 {
		return false
	} else if grid[setOffX][setOffY].GetSectorName() != hexSectors.DangerousName {
		return false
	}

	return true
}