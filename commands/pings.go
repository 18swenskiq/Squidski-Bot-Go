package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) ChangePingState(session *discordgo.Session, message *discordgo.MessageCreate, pingsRole string) {
	checkIfSubbedToPings := false
	pingsUser, err := session.GuildMember(message.GuildID, message.Author.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range pingsUser.Roles {
		if i == pingsRole {
			checkIfSubbedToPings = true
		}
	}
	if checkIfSubbedToPings {
		err := session.GuildMemberRoleRemove(message.GuildID, message.Author.ID, pingsRole)
		if err != nil {
			fmt.Println(err)
		}
		session.ChannelMessageSend(message.ChannelID, "You have been unsubscribed from the 'pings' role!")
	} else {
		err := session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, pingsRole)
		if err != nil {
			fmt.Println(err)
		}
		session.ChannelMessageSend(message.ChannelID, "You are now subscribed to the 'pings' role!")
	}
}
