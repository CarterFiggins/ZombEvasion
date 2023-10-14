package game

import (
	"fmt"
	"math/rand"
	"time"

	"infection/models"
	"infection/hexagonGrid/hexSectors"
	"infection/hexagonGrid"
	"github.com/bwmarrin/discordgo"
)

func MovedOnSectorMessages(interaction *discordgo.InteractionCreate, sectorName, hexName string) (string, string, error) {
	if (sectorName == hexSectors.SafeHouseName) {
		userMessage := "You made it to the Save House!"
		turnMessage := fmt.Sprintf("%v has made it to the save house!", interaction.Interaction.Member.User.Mention())
		return turnMessage, userMessage, nil
	}

	turnMessage := "Silence No alarm"
	userMessage := fmt.Sprintf("You moved to a %s located at: %s\n Silence. No alarms where set off", sectorName, hexName)

	if (sectorName == hexSectors.DangerousName) {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(10)
		// 40% chance green
		if randNum >= 0 && randNum <= 3 {
			userMessage =fmt.Sprintf("You moved to a %s located at: %s\n You get to set off an alarm in another sector. Use `/set-off-alarm` to pick location", sectorName, hexName)
			turnMessage = ""
			mongoUser, err := models.FindUser(interaction, nil)
			if err != nil {
				return "", "", err
			}
			err = mongoUser.UpdateCanSetOffAlarm(true)
			if err != nil {
				return "", "", err
			}
		}
		// 40% chance red
		if randNum >= 4 && randNum <= 7 {
			userMessage = fmt.Sprintf("You moved to a %s located at: %s\n The Alarm was set off!", sectorName, hexName)
			turnMessage = fmt.Sprintf("ALERT! Alarm set off at %s", hexName)
		}
		// 20% change silence
		if randNum >= 8 && randNum <= 9 {
			userMessage = fmt.Sprintf("You moved to a %s located at: %s\n Silence. No alarms where set off", sectorName, hexName)
		}
	}

	return turnMessage, userMessage, nil
}

func CanUserMoveHere(discord *discordgo.Session, interaction *discordgo.InteractionCreate, moveX int, moveY int, mongoUser *models.MongoUser) (*string, error) {
	if (!mongoUser.CanMove) {
		message := "You have already moved this turn. Use `/end-turn` to start the next players turn"
		return &message, nil
	}

	moveHexName := hexSectors.GetHexName(moveX, moveY)

	sectorsToMove := GetMoveSectors(mongoUser)

	if (moveX == mongoUser.Col && moveY == mongoUser.Row) {
		message := fmt.Sprintf("You can't move to your current position: %s.\nAvailable sectors to move: %v", mongoUser.Location.GetHexName(), sectorsToMove)
		return &message, nil
	}

	for _, sector := range sectorsToMove {
		if sector == hexSectors.GetHexName(moveX, moveY) {
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

	travelSectors(mongoUser, &sectorSlice, 0, mongoUser.MaxMoves)

	return sectorSlice[1:]
}

func travelSectors(mongoUser *models.MongoUser, sectorSlice *[]string, depth, limit int) {
	if depth == limit {
		return
	}

	location := mongoUser.Location
	up := location.Col % 2 == 0
	
	if up {
		for i := location.Col - 1; i < location.Col + 2; i++ {
			for j := location.Row - 1; j < location.Row + 1; j++ {
				if canMoveHere(i, j, mongoUser.Role) {
					location := &hexSectors.Location{Col: i, Row: j}
					addSector(location, sectorSlice)
					travelSectors(mongoUser, sectorSlice, depth + 1, limit)
				}
			}
		}
		if canMoveHere(location.Col, location.Row + 1, mongoUser.Role) {
			location := &hexSectors.Location{Col: location.Col, Row: location.Row - 1}
			addSector(location, sectorSlice)
			travelSectors(mongoUser, sectorSlice, depth + 1, limit)
			
		}
	} else {
		for i := location.Col - 1; i < location.Col + 2; i++ {
			for j := location.Row; j < location.Row + 2; j++ {
				if canMoveHere(i, j, mongoUser.Role) {
					location := &hexSectors.Location{Col: i, Row: j}
					addSector(location, sectorSlice)
					travelSectors(mongoUser, sectorSlice, depth + 1, limit)
					
				}
			}
		}
		if canMoveHere(location.Col, location.Row - 1, mongoUser.Role) {
			location := &hexSectors.Location{Col: location.Col, Row: location.Row - 1}
			addSector(location, sectorSlice)
			travelSectors(mongoUser, sectorSlice, depth + 1, limit)
		}
	}
	return
}

func canMoveHere(col, row int, userRole string) bool {
	grid := hexagonGrid.Board.Grid
	if col >= len(grid) || col < 0 {
		return false
	} else if row >= len(grid[col]) || row < 0 {
		return false
	} else if !grid[col][row].CanMoveHere() {
		return false
	} else if userRole == models.Zombie && grid[col][row].GetSectorName() == hexSectors.SafeHouseName {
		return false
	}

	return true
}

func addSector(location *hexSectors.Location, sectorSlice *[]string) bool {
	hexName := location.GetHexName()

	for _, name := range *sectorSlice {
		if name == hexName {
			return false
		}
	}

	*sectorSlice = append(*sectorSlice, hexName)

	return true
}
