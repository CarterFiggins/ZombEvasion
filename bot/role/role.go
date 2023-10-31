package role

import (
	"log"
	"errors"

	"github.com/bwmarrin/discordgo"
)

const (
	WaitingForNextGame string = "WaitingForNextGame"
	Admin = "Admin"
	InGame = "InGame"
	Bot = "Infection"
)

func isAdmin(discord *discordgo.Session, interaction *discordgo.InteractionCreate) (bool, error) {
	roleMap, err := CreateRoleMap(discord, interaction)
	if err != nil {
		return false, err
	}
	adminRole, ok := roleMap[Admin]
	if (!ok) {
		return false, errors.New("role not found. Might need to run `/setup-server` to add the roles.")
	}
	roles := interaction.Interaction.Member.Roles
	for _, roleID := range roles {
		if roleID == adminRole.ID {
			return true, nil
		}
	}
	return false, nil
}

func UserHasOneRole(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roles []string) (bool, error) {
	hasRoleMap := make(map[string]bool)
	for _, role := range roles {
		hasRoleMap[role] = false
	}
	
	roleMap, err := CreateRoleMap(discord, interaction)
	if err != nil {
		return false, err
	}

	userRoles := interaction.Interaction.Member.Roles
	for _, roleName := range roles {
		role, ok := roleMap[roleName]
		if !ok {
			continue
		}
		for _, userRoleID := range userRoles {
			if userRoleID == role.ID {
				hasRoleMap[roleName] = true
			}
		}
	}

	for _, isRole := range hasRoleMap {
		if isRole {
			return true, nil
		}
	}
	return false, nil
}

func SetUpRoles(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error{
	log.Println("SetUpRoles on discord")
	guildID := interaction.Interaction.GuildID
	roleMap, err := CreateRoleMap(discord, interaction)
	if err != nil {
		return err
	}

	if _, ok := roleMap[WaitingForNextGame]; !ok {
		err = makeRole(discord, guildID, WaitingForNextGame, 12285184)
		if err != nil {
			return err
		}
	}
	if _, ok := roleMap[InGame]; !ok {
		err = makeRole(discord, guildID, InGame, 44853)
		if err != nil {
			return err
		}
	}
	if _, ok := roleMap[Admin]; !ok {
		err = makeRole(discord, guildID, Admin, 34223)
		if err != nil {
			return err
		}
	}

	return nil
}

func AddRole(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roleName string) error {
	role, err := GetRole(discord, interaction, roleName)
	if err != nil {
		return err
	}

	return discord.GuildMemberRoleAdd(
		interaction.Interaction.GuildID,
		interaction.Interaction.Member.User.ID,
		role.ID,
	)
}

func RemoveRole(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roleName string) error {
	role, err := GetRole(discord, interaction, roleName)
	if err != nil {
		return err
	}

	return discord.GuildMemberRoleRemove(
		interaction.Interaction.GuildID,
		interaction.Interaction.Member.User.ID,
		role.ID,
	)
}

func AddRoleToUsers(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roleName string, users []*discordgo.User) error {
	role, err := GetRole(discord, interaction, roleName)
	if err != nil {
		return err
	}

	for _, user := range users {
		err = discord.GuildMemberRoleAdd(
			interaction.Interaction.GuildID,
			user.ID,
			role.ID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveRoleToUsers(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roleName string, users []*discordgo.User) error {
	role, err := GetRole(discord, interaction, roleName)
	if err != nil {
		return err
	}

	for _, user := range users {
		err = discord.GuildMemberRoleRemove(
			interaction.Interaction.GuildID,
			user.ID,
			role.ID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}


func GetRole(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roleName string) (*discordgo.Role, error) {
	roleMap, err := CreateRoleMap(discord, interaction)
	if err != nil {
		return nil, err
	}

	role, ok := roleMap[roleName]
	if (!ok) {
		return nil, errors.New(roleName + " not found. Might need to run `/setup-server` to add the roles.")
	}

	return role, nil
}

func makeRole(discord *discordgo.Session, guildId string, name string, color int) error {
	mentionable := true
	_, err := discord.GuildRoleCreate(
		guildId,
		&discordgo.RoleParams{
			Name: name,
			Color: &color,
			Mentionable: &mentionable,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func CreateRoleMap(discord *discordgo.Session, interaction *discordgo.InteractionCreate) (map[string]*discordgo.Role, error){
	roles, err := discord.GuildRoles(interaction.Interaction.GuildID)
	roleMap := make(map[string]*discordgo.Role)
	if (err != nil) {
		return roleMap, err
	}
	for _, role := range roles {
		roleMap[role.Name] = role
	}
	return roleMap, nil
}

func GetDiscordUsersFromRole(discord *discordgo.Session, interaction *discordgo.InteractionCreate, roleName string) ([]*discordgo.User, error) {
	members, err := discord.GuildMembers(interaction.Interaction.GuildID, "1", 1000)
	if err != nil {
		return nil, err
	}

	discordRole, err := GetRole(discord, interaction, roleName)
	if err != nil {
		return nil, err
	}
	
	
	var users []*discordgo.User
	for _, member := range members {
		for _, roleID := range member.Roles {
			if roleID == discordRole.ID {
				users = append(users, member.User)
				break
			}
		}
	}
	return users, nil
}

func GetAllDiscordUsers(discord *discordgo.Session, interaction *discordgo.InteractionCreate) ([]*discordgo.User, error) {
	members, err := discord.GuildMembers(interaction.Interaction.GuildID, "1", 1000)
	if err != nil {
		return nil, err
	}

	var users []*discordgo.User
	for _, member := range members {
		users = append(users, member.User)
	}
	return users, nil
}

func MoveRoleToInGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate, usersInQueue []*discordgo.User) error {
	err := AddRoleToUsers(discord, interaction, InGame, usersInQueue)
	if err != nil {
		return err
	}
	err = RemoveRoleToUsers(discord, interaction, WaitingForNextGame, usersInQueue)
	if err != nil {
		return err
	}

	return nil
}

func RemoveInGameRole(discord *discordgo.Session, interaction *discordgo.InteractionCreate, users []*discordgo.User) error {
	err := RemoveRoleToUsers(discord, interaction, InGame, users)
	if err != nil {
		return err
	}
	
	return nil
}