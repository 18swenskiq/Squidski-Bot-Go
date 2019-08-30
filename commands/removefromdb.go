package commands

import (
	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) RemoveFromDB(session *discordgo.Session, message *discordgo.MessageCreate, messageArray []string, ownerID string) {
	if message.Author.ID != ownerID {
		session.ChannelMessageSend(message.ChannelID, "Only the owner of this server can add to the database...")
		return
	}
	if len(messageArray) < 3 {
		session.ChannelMessageSend(message.ChannelID, "Not enough parameters were given to add to the database!")
		return
	}
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	database.DeleteKey(messageArray[1], messageArray[2])
	session.ChannelMessageSend(message.ChannelID, "Key "+messageArray[2]+" was deleted from the bucket "+messageArray[1])
}
