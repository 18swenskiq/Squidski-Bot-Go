package commands

import (
	"strings"

	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) AddToDB(session *discordgo.Session, message *discordgo.MessageCreate, messageArray []string, ownerID string) {
	if message.Author.ID != ownerID {
		session.ChannelMessageSend(message.ChannelID, "Only the owner of this server can add to the database...")
		return
	}
	if len(messageArray) < 4 {
		session.ChannelMessageSend(message.ChannelID, "Not enough parameters were given to add to the database!")
		return
	}
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	database.WriteToDB(messageArray[1], messageArray[2], strings.Join(messageArray[3:], " "))
	session.ChannelMessageSend(message.ChannelID, "KeyName \"**"+messageArray[2]+"**\" with value of \"**"+strings.Join(messageArray[3:], " ")+"**\" was written to bucket \"**"+messageArray[1]+"**\".")
}
