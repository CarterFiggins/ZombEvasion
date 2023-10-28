package game

import (
	"fmt"

	"infection/models"
	"infection/hexagonGrid/hexSectors"
	"infection/hexagonGrid"
	"infection/bot/channel"
	"github.com/bwmarrin/discordgo"
)

func AttackSector(discord *discordgo.Session, guildID string, mongoUser *models.MongoUser, attackX, attackY int) ([]string, bool, error) {
	usersAttacked, err := models.FindUsersAtLocation(guildID, attackX, attackY)
	var usersAttackedRoles []string
	zombieUpgrade := false

	for _, user := range usersAttacked {
		if (user.DiscordUserID == mongoUser.DiscordUserID) {
			// Don't attack self
			continue
		}
		usersAttackedRoles = append(usersAttackedRoles, user.Role)
		zombieSectorCol := hexagonGrid.Board.ZombieSector.Col
		zombieSectorRow := hexagonGrid.Board.ZombieSector.Row

		if (user.Role == models.Human) {
			if (mongoUser.MaxMoves == 2) {
				zombieUpgrade = true
				mongoUser.UpgradeUsersMaxMoves(3)
			}
			attackedMessage := fmt.Sprintf("You have been bitten by a zombie! You have Respawned as a zombie at %s", hexSectors.GetHexName(zombieSectorCol, zombieSectorRow))
			if err = channel.SendUserMessage(discord, user.DiscordUserID, attackedMessage); err != nil {
				return usersAttackedRoles, false, err
			}

			if err = user.TurnIntoZombie(); err != nil {
				return usersAttackedRoles, false, err
			}

			if err = CheckGame(discord, guildID); err != nil {
				return usersAttackedRoles, false, err
			} 
			
		} else if (user.Role == models.Zombie) {
			attackedMessage := fmt.Sprintf("A zombie mistaken you as a human and attacked you! You have Respawned at %s", hexSectors.GetHexName(zombieSectorCol, zombieSectorRow))
			if err = channel.SendUserMessage(discord, user.DiscordUserID, attackedMessage); err != nil {
				return usersAttackedRoles, false, err
			}
		}
		if err = user.RespawnUser(zombieSectorCol, zombieSectorRow); err != nil {
			return usersAttackedRoles, false, err
		}
	}
	return usersAttackedRoles, zombieUpgrade, nil
}
