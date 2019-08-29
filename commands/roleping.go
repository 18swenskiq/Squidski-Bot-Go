package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) PingRole(session *discordgo.Session, message *discordgo.MessageCreate, adminRole string, pingsRole string) {
	adminUser, err := session.GuildMember(message.GuildID, message.Author.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	checkForAdminFlag := false
	for _, i := range adminUser.Roles {
		if i == adminRole {
			checkForAdminFlag = true
		}
	}
	if !checkForAdminFlag {
		session.ChannelMessageSend(message.ChannelID, "You must be an administrator to use this command...")
		return
	}
	t, err := session.GuildRoles(message.GuildID)
	for _, j := range t {
		if j.ID == pingsRole {
			session.GuildRoleEdit(message.GuildID, pingsRole, j.Name, j.Color, j.Hoist, j.Permissions, true)
			break
		}
	}
	session.ChannelMessageSend(message.ChannelID, "<@&"+pingsRole+">\n*To unsubscribe from pings, type >pings")
	for _, j := range t {
		if j.ID == pingsRole {
			session.GuildRoleEdit(message.GuildID, pingsRole, j.Name, j.Color, j.Hoist, j.Permissions, false)
			break
		}
	}
}
