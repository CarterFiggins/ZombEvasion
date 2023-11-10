package game

import (
	"fmt"

	"infection/hexagonGrid"
	"infection/hexagonGrid/hexSectors"
)

func CanSetOffAlarmHere(setOffX, setOffY int, guildID string) bool {
	grid, err := hexagonGrid.GetBoard(guildID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if setOffX >= len(grid) || setOffX < 0 {
		return false
	} else if setOffY >= len(grid[setOffX]) || setOffY < 0 {
		return false
	} else if grid[setOffX][setOffY].GetSectorName() != hexSectors.DangerousName {
		return false
	}

	return true
}