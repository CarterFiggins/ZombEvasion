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

type depthSector struct {
	Depth int
	Name string
}

type SectorsToMove struct {
	DepthSectors []depthSector
}

func (s *SectorsToMove) GetAllSectors() []string {
	var sectors []string
	
	for _, depthSector := range s.DepthSectors {
		sectors = append(sectors, depthSector.Name)
	}
	return sectors
}

func (s *SectorsToMove) GetDepthSectors(depth int) []string {
	var sectors []string
	
	for _, depthSector := range s.DepthSectors {
		if depthSector.Depth == depth {
			sectors = append(sectors, depthSector.Name)
		}
	}
	return sectors
}

func MovedOnSectorMessages(discord *discordgo.Session, interaction *discordgo.InteractionCreate, sectorName, hexName string) (string, string, error) {
	mongoUser, err := models.FindUser(interaction)
	if err != nil {
		return "", "", err
	}
	
	if (sectorName == hexSectors.SafeHouseName) {
		userMessage := "You made it to the Save House!"
		turnMessage := fmt.Sprintf("%v has made it to the save house!", interaction.Interaction.Member.User.Mention())

		if err = mongoUser.EnterSafeHouse(); err != nil {
			return "", "", err
		}

		if err = CheckGame(discord, interaction); err != nil {
			return "", "", err
		}

		NextTurn(discord, interaction, mongoUser)

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
			err = mongoUser.UpdateCanSetOffAlarm(true)
			if err != nil {
				return "", "", err
			}
		}
		// 40% chance red
		if randNum >= 4 && randNum <= 7 {
			userMessage = fmt.Sprintf("You moved to a %s located at: %s\n The Alarm was set off!", sectorName, hexName)
			turnMessage = fmt.Sprintf("ALERT! Alarm set off at %s", hexName)
			NextTurn(discord, interaction, mongoUser)
		}
		// 20% change silence
		if randNum >= 8 && randNum <= 9 {
			userMessage = fmt.Sprintf("You moved to a %s located at: %s\n Silence. No alarms where set off", sectorName, hexName)
			NextTurn(discord, interaction, mongoUser)
		}
	} else {
		NextTurn(discord, interaction, mongoUser)
	}

	return turnMessage, userMessage, nil
}

func CanUserMoveHere(discord *discordgo.Session, interaction *discordgo.InteractionCreate, moveX int, moveY int, mongoUser *models.MongoUser) (*string, error) {

	moveHexName := hexSectors.GetHexName(moveX, moveY)

	sectorsToMove := GetMoveSectors(mongoUser)

	if (moveX == mongoUser.Col && moveY == mongoUser.Row) {
		message := fmt.Sprintf("You can't move to your current position: %s.\nAvailable sectors to move: %v", mongoUser.Location.GetHexName(), sectorsToMove.GetAllSectors())
		return &message, nil
	}

	for _, sector := range sectorsToMove.GetAllSectors() {
		if sector == hexSectors.GetHexName(moveX, moveY) {
			return nil, nil
		}
	}

	message := fmt.Sprintf("You can't move to position: %s.\nCurrent position: %s\nAvailable sectors to move: %v", moveHexName, mongoUser.Location.GetHexName(), sectorsToMove.GetAllSectors())
	return &message, nil
}

func GetMoveSectors(mongoUser *models.MongoUser) *SectorsToMove {
	sectorsToMove := &SectorsToMove{
		DepthSectors: []depthSector{
			{
				Depth: 0,
				Name: mongoUser.Location.GetHexName(),
			},
		},
	}

	travelSectors(mongoUser.Location, mongoUser.Role, sectorsToMove, 1, mongoUser.MaxMoves)

	// remove users current position
	sectorsToMove.DepthSectors = sectorsToMove.DepthSectors[1:]
	return sectorsToMove
}

func travelSectors(location *hexSectors.Location, userRole string, sectorsToMove *SectorsToMove, depth, limit int) {
	if depth > limit {
		return
	}

	up := location.Col % 2 == 0
	
	if up {
		for i := location.Col - 1; i < location.Col + 2; i++ {
			for j := location.Row - 1; j < location.Row + 1; j++ {
				if canMoveHere(i, j, userRole) {
					newLocation := &hexSectors.Location{Col: i, Row: j}
					addSector(newLocation, depth, sectorsToMove)
					travelSectors(newLocation, userRole, sectorsToMove, depth + 1, limit)
				}
			}
		}
		if canMoveHere(location.Col, location.Row + 1, userRole) {
			newLocation := &hexSectors.Location{Col: location.Col, Row: location.Row + 1}
			addSector(newLocation, depth, sectorsToMove)
			travelSectors(newLocation, userRole, sectorsToMove, depth + 1, limit)
			
		}
	} else {
		for i := location.Col - 1; i < location.Col + 2; i++ {
			for j := location.Row; j < location.Row + 2; j++ {
				if canMoveHere(i, j, userRole) {
					newLocation := &hexSectors.Location{Col: i, Row: j}
					addSector(newLocation, depth, sectorsToMove)
					travelSectors(newLocation, userRole, sectorsToMove, depth + 1, limit)
					
				}
			}
		}
		if canMoveHere(location.Col, location.Row - 1, userRole) {
			newLocation := &hexSectors.Location{Col: location.Col, Row: location.Row - 1}
			addSector(newLocation, depth, sectorsToMove)
			travelSectors(newLocation, userRole, sectorsToMove, depth + 1, limit)
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

func addSector(location *hexSectors.Location, depth int, sectorsToMove *SectorsToMove) bool {
	hexName := location.GetHexName()

	for _, sector := range sectorsToMove.DepthSectors {
		if sector.Name == hexName {
			return false
		}
	}

	sectorsToMove.DepthSectors = append(sectorsToMove.DepthSectors, depthSector{
		Name: hexName,
		Depth: depth,
	})

	return true
}

func RespondWithLocationButtons(discord *discordgo.Session, interaction *discordgo.InteractionCreate, mongoUser *models.MongoUser, guildID string) {
	sectorsToMove := GetMoveSectors(mongoUser)
	allSectors := sectorsToMove.GetAllSectors()
	var components []discordgo.MessageComponent
	var buttons []discordgo.MessageComponent
	content := "Select where to move"
	if mongoUser.IsAttacking {
		content = "Select where to attack"
	}
	
	if len(allSectors) > 24 {
		content = "Select how far you want to move"
		if mongoUser.IsAttacking {
			content = "Select how far you want to attack"
		}
		for move := 1; move <= mongoUser.MaxMoves; move++ {
			buttons = append(buttons, discordgo.Button{
				Label: fmt.Sprintf("Move %d sector(s)", move),
				CustomID: fmt.Sprintf("max-move_%s_%d", guildID, move),
			})
		}
		components = append(components, discordgo.ActionsRow{
			Components: buttons,
		})
	} else {
		for index, sectorName := range allSectors {
			buttons = append(buttons, discordgo.Button{
				Label: sectorName,
				CustomID: fmt.Sprintf("move-user_%s_%s", guildID, sectorName),
			})
			if (index + 1) % 5 == 0 {
				components = append(components, discordgo.ActionsRow{
					Components: buttons,
				})
				buttons = []discordgo.MessageComponent{}
			}
		}
		if len(buttons) > 0 {
			components = append(components, discordgo.ActionsRow{
				Components: buttons,
			})
		}
	}

	response := &discordgo.WebhookEdit{
		Content: &content,
		Components: &components,
	}

	discord.InteractionResponseEdit(interaction.Interaction, response)
}
