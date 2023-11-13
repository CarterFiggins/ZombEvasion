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
	DepthSectors []*depthSector
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

type TravelData struct {
	MongoUser *models.MongoUser
	Location *hexSectors.Location
	Grid [][]hexagonGrid.Hex
	SectorsToMove *SectorsToMove
	Depth int
}

type TravelPoint struct {
	Data *TravelData
	Location *hexSectors.Location
}

func MovedOnSectorMessages(discord *discordgo.Session, interaction *discordgo.InteractionCreate, sectorName, hexName, guildID string) (string, string, error) {
	mongoUser, err := models.FindUserByIDs(interaction, nil, &guildID)
	if err != nil {
		return "", "", err
	}

	discordUser := interaction.Interaction.User
	if discordUser == nil {
		discordUser = interaction.Interaction.Member.User
	}
	
	if (sectorName == hexSectors.SafeHouseName) {
		userMessage := "You made it to the Safe House!"
		turnMessage := fmt.Sprintf("%v has made it to the safe house!", discordUser.Mention())

		if err = mongoUser.EnterSafeHouse(); err != nil {
			return "", "", err
		}

		if err = CheckGame(discord, guildID); err != nil {
			return "", "", err
		}
		return turnMessage, userMessage, nil
	}

	turnMessage := "Silence No alarm"
	userMessage := fmt.Sprintf("You moved to a %s located at: %s\n Silence. No alarms where set off", sectorName, hexName)

	if (sectorName == hexSectors.DangerousName) {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(10)
		// 40% chance green
		if randNum >= 0 && randNum <= 3 {
			userMessage =fmt.Sprintf("You moved to a %s located at: %s\n You get to set off an alarm in another sector. Type a sector in the text input. e.g A01", sectorName, hexName)
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
		}
		// 20% change silence
		if randNum >= 8 && randNum <= 9 {
			userMessage = fmt.Sprintf("You moved to a %s located at: %s\n Silence. No alarms where set off", sectorName, hexName)
		}
	}

	return turnMessage, userMessage, nil
}

func CanUserMoveHere(discord *discordgo.Session, interaction *discordgo.InteractionCreate, moveHexName, guildID string, mongoUser *models.MongoUser) (*string, error) {
	sectorsToMove, err := GetMoveSectors(mongoUser, guildID)
	if err != nil {
		return nil, err
	}
	
	action := "move to"
	if mongoUser.IsAttacking {
		action = "attack"
	}

	if (moveHexName == mongoUser.Location.GetHexName()) {
		message := fmt.Sprintf("You can't %s your current position: %s.\nAvailable sectors to move: %v", action, mongoUser.Location.GetHexName(), sectorsToMove.GetAllSectors())
		return &message, nil
	}

	for _, sector := range sectorsToMove.GetAllSectors() {
		if sector == moveHexName {
			return nil, nil
		}
	}

	message := fmt.Sprintf("You can't %s position: %s.\nCurrent position: %s\nAvailable sectors to move: %v", action, moveHexName, mongoUser.Location.GetHexName(), sectorsToMove.GetAllSectors())
	return &message, nil
}

func GetMoveSectors(mongoUser *models.MongoUser, guildID string) (*SectorsToMove, error) {
	sectorsToMove := &SectorsToMove{
		DepthSectors: []*depthSector{
			{
				Depth: 0,
				Name: mongoUser.Location.GetHexName(),
			},
		},
	}

	grid, err := hexagonGrid.GetBoard(guildID)
	if err != nil {
		return nil, err
	}

	travelData := &TravelData{
		MongoUser: mongoUser,
		Grid: grid,
		SectorsToMove: sectorsToMove,
	}
	
	travelSectors(travelData, mongoUser.Location, 0)


	// remove users current position
	travelData.SectorsToMove.DepthSectors = travelData.SectorsToMove.DepthSectors[1:]
	return travelData.SectorsToMove, nil
}

func travelSectors(travelData *TravelData, location *hexSectors.Location, depth int) {
	depth += 1
	if depth > travelData.MongoUser.MaxMoves {
		return
	}

	up := location.Col % 2 == 0
	travelSlice := []*TravelPoint{}
	
	if up {
		for i := location.Col - 1; i < location.Col + 2; i++ {
			for j := location.Row - 1; j < location.Row + 1; j++ {
				if canMoveHere(i, j, travelData, depth) {
					newLocation := &hexSectors.Location{Col: i, Row: j}
					travelSlice = append(travelSlice, &TravelPoint{Data: travelData, Location: newLocation})
				}
			}
		}
		if canMoveHere(location.Col, location.Row + 1, travelData, depth) {
			newLocation := &hexSectors.Location{Col: location.Col, Row: location.Row + 1}
			travelSlice = append(travelSlice, &TravelPoint{Data: travelData, Location: newLocation})
			
		}
	} else {
		for i := location.Col - 1; i < location.Col + 2; i++ {
			for j := location.Row; j < location.Row + 2; j++ {
				if canMoveHere(i, j, travelData, depth) {
					newLocation := &hexSectors.Location{Col: i, Row: j}
					travelSlice = append(travelSlice, &TravelPoint{Data: travelData, Location: newLocation})
				}
			}
		}
		if canMoveHere(location.Col, location.Row - 1, travelData, depth) {
			newLocation := &hexSectors.Location{Col: location.Col, Row: location.Row - 1}
			travelSlice = append(travelSlice, &TravelPoint{Data: travelData, Location: newLocation})
		}
	}

	for _, travelPoint := range travelSlice {
		addSector(travelPoint.Data, travelPoint.Location, depth)
		travelSectors(travelPoint.Data, travelPoint.Location, depth)
	}
	return
}

func canMoveHere(col, row int, travelData *TravelData, depth int) bool {
	if col >= len(travelData.Grid) || col < 0 {
		return false
	} else if row >= len(travelData.Grid[col]) || row < 0 {
		return false
	} else if !travelData.Grid[col][row].CanMoveHere() {
		return false
	} else if travelData.MongoUser.Role == models.Zombie && travelData.Grid[col][row].GetSectorName() == hexSectors.SafeHouseName {
		return false
	}
	location := &hexSectors.Location{Col: col, Row: row}
	for _, move := range travelData.SectorsToMove.DepthSectors {
		if move.Name == location.GetHexName() && depth >= move.Depth {
			return false
		}
	}

	return true
}

func addSector(travelData *TravelData, location *hexSectors.Location, depth int) {
	hexName := location.GetHexName()

	for _, sector := range travelData.SectorsToMove.DepthSectors {
		if sector.Name == hexName {
			if sector.Depth > depth {
				sector.Depth = depth
				return
			} else {
				return
			}
		}
	}

	travelData.SectorsToMove.DepthSectors = append(travelData.SectorsToMove.DepthSectors, &depthSector{
		Name: hexName,
		Depth: depth,
	})
}

func makeMoveButtons(sectors []string, guildID string) []discordgo.MessageComponent {
	var components []discordgo.MessageComponent
	var buttons []discordgo.MessageComponent

	for index, sectorName := range sectors {
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

	return components
}

func RespondWithDistanceButtons(discord *discordgo.Session, interaction *discordgo.InteractionCreate, mongoUser *models.MongoUser, guildID string, moveDistance int) error {
	sectorsToMove, err := GetMoveSectors(mongoUser, guildID)
	if err != nil {
		return err
	}
	
	sectorsFromDistance := sectorsToMove.GetDepthSectors(moveDistance)
	content := "Select where to move"
	if mongoUser.IsAttacking {
		content = "Select where to attack"
	}

	/*
		If the Distance is above 4 it can break because you can't have more than
		25 buttons in a message. If Distance is at 5, player could have 30 possible
		sectors to move to that are 5 sectors away. Another way to solve this is to
		send multiple messages each have 25 buttons or lower.
	*/
	components := makeMoveButtons(sectorsFromDistance, guildID)

	response := &discordgo.WebhookEdit{
		Content: &content,
		Components: &components,
	}

	_, err = discord.InteractionResponseEdit(interaction.Interaction, response)
	return err
}

func RespondWithLocationButtons(discord *discordgo.Session, interaction *discordgo.InteractionCreate, mongoUser *models.MongoUser, guildID string) error {
	sectorsToMove, err := GetMoveSectors(mongoUser, guildID)
	if err != nil {
		return err
	}
	allSectors := sectorsToMove.GetAllSectors()
	var components []discordgo.MessageComponent
	var buttons []discordgo.MessageComponent
	content := "Select where to move"
	if mongoUser.IsAttacking {
		content = "Select where to attack"
	}
	
	// Can't send a message that has more than 25 buttons
	if len(allSectors) > 25 {
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
		components = makeMoveButtons(allSectors, guildID)
	}

	response := &discordgo.WebhookEdit{
		Content: &content,
		Components: &components,
	}

	_, err = discord.InteractionResponseEdit(interaction.Interaction, response)
	return err
}
