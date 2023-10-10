package game

import (
	"fmt"

	"infection/models"
	"infection/hexagonGrid/hexTypes"
	"infection/hexagonGrid"
	"github.com/bwmarrin/discordgo"
)

func CanUserMoveHere(discord *discordgo.Session, interaction *discordgo.InteractionCreate, moveX int, moveY int) (*string, error) {
	mongoUser, err := models.FindUser(interaction.Interaction.GuildID, interaction.Interaction.Member.User.ID)
	if err != nil {
		return nil, err
	}

	if (mongoUser.Moved) {
		message := "You have already moved this turn use `/end-turn` to pass it to the next player"
		return &message, nil
	}

	moveHexName := hexTypes.GetHexName(moveX, moveY)

	sectorsToMove := GetMoveSectors(mongoUser)

	if (moveX == mongoUser.Col && moveY == mongoUser.Row) {
		message := fmt.Sprintf("You can't move to your current position: %s.\nAvailable sectors to move: %v", mongoUser.Location.GetHexName(), sectorsToMove)
		return &message, nil
	}

	for _, sector := range sectorsToMove {
		if sector == hexTypes.GetHexName(moveX, moveY) {
			return nil, nil
		}
	}

	message := fmt.Sprintf("You can't move to position: %s.\nCurrent position: %s\nAvailable sectors to move: %v", moveHexName, mongoUser.Location.GetHexName(), sectorsToMove)
	return &message, nil
}

func GetMoveSectors(mongoUser *models.MongoUser) []string {
	sectorSlice := []string{
		mongoUser.Location.GetHexName(),
	}

	travelSectors(mongoUser.Location, &sectorSlice, 0, mongoUser.MaxMoves)

	return sectorSlice[1:]
}

func travelSectors(location *hexTypes.Location, sectorSlice *[]string, depth, limit int) {
	if depth == limit {
		return
	}
	
	up := location.Col % 2 == 0
	
	if up {
		for i := location.Col - 1; i < location.Col + 2; i++ {
			for j := location.Row - 1; j < location.Row + 1; j++ {
				if canMoveHere(i, j) {
					location := &hexTypes.Location{Col: i, Row: j}
					addSector(location, sectorSlice)
					travelSectors(location, sectorSlice, depth + 1, limit)
				}
			}
		}
		if canMoveHere(location.Col, location.Row + 1) {
			location := &hexTypes.Location{Col: location.Col, Row: location.Row - 1}
			addSector(location, sectorSlice)
			travelSectors(location, sectorSlice, depth + 1, limit)
			
		}
	} else {
		for i := location.Col - 1; i < location.Col + 2; i++ {
			for j := location.Row; j < location.Row + 2; j++ {
				if canMoveHere(i, j) {
					location := &hexTypes.Location{Col: i, Row: j}
					addSector(location, sectorSlice)
					travelSectors(location, sectorSlice, depth + 1, limit)
					
				}
			}
		}
		if canMoveHere(location.Col, location.Row - 1) {
			location := &hexTypes.Location{Col: location.Col, Row: location.Row - 1}
			addSector(location, sectorSlice)
			travelSectors(location, sectorSlice, depth + 1, limit)
		}
	}
	return
}

func canMoveHere(col, row int) bool {
	grid := hexagonGrid.Board.Grid
	if col >= len(grid) || col < 0 {
		return false
	} else if row >= len(grid[col]) || row < 0 {
		return false
	} else if !grid[col][row].CanMoveHere() {
		return false
	}

	return true
}

func addSector(location *hexTypes.Location, sectorSlice *[]string) bool {
	hexName := location.GetHexName()

	for _, name := range *sectorSlice {
		if name == hexName {
			return false
		}
	}

	*sectorSlice = append(*sectorSlice, hexName)

	return true
}
