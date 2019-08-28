package commands

import (
	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) ViewMutes(session *discordgo.Session, message *discordgo.MessageCreate) {

	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	names, values := database.IterateOverKeysInBucketReturnBoth("MutedUsers")
	if len(names) == 0 {
		session.ChannelMessageSend(message.ChannelID, "No users are currently muted (according to the database)")
	} else {
		stringToPrint := "\n"
		for i := range names {
			stringToPrint += "<@" + names[i] + "> is muted until " + values[i] + "\n"
		}
		session.ChannelMessageSend(message.ChannelID, stringToPrint)
	}
}
