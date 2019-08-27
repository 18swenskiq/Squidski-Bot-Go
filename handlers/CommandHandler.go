package handlers

import (
	"strings"

	commands "../commands"
	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	session discordgo.Session
	message discordgo.MessageCreate
}

func (c *CommandHandler) ExecuteCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	messageArray := strings.Split(message.Content, " ")
	messageArray[0] = strings.ToLower(messageArray[0][1:])

	if messageArray[0] == "help" {
		var helpEmbed *commands.GenericCommand
		helpEmbed = new(commands.GenericCommand)
		helpEmbed.BuildHelpEmbed(session, message)
	}
}
