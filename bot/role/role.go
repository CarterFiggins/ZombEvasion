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
)

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
	if _, ok := roleMap[Admin]; !ok {
		err = makeRole(discord, guildID, Admin, 34223)
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

func MoveRolesInGame(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	members, err := discord.GuildMembers(interaction.Interaction.GuildID, "1", 1000)
	if err != nil {
		return err
	}

	var usersInQueue []*discordgo.User
	discordRole, err := GetRole(discord, interaction, WaitingForNextGame)

	if err != nil {
		return err
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if roleID == discordRole.ID {
				usersInQueue = append(usersInQueue, member.User)
				break
			}
		}
	}


	err = AddRoleToUsers(discord, interaction, InGame, usersInQueue)
	if err != nil {
		return err
	}
	err = RemoveRoleToUsers(discord, interaction, WaitingForNextGame, usersInQueue)
	if err != nil {
		return err
	}

	return nil
}

func RemoveAllInGameRoles(discord *discordgo.Session, interaction *discordgo.InteractionCreate) error {
	members, err := discord.GuildMembers(interaction.Interaction.GuildID, "1", 1000)
	if err != nil {
		return err
	}

	var users []*discordgo.User
	for _, member := range members {
		users = append(users, member.User)
	}

	err = RemoveRoleToUsers(discord, interaction, InGame, users)
	if err != nil {
		return err
	}
	
	return nil
}