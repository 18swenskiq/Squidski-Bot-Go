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

func (c *CommandHandler) ExecuteCommand(session *discordgo.Session, message *discordgo.MessageCreate, adminRole string, mutedRole string) {
	messageArray := strings.Split(message.Content, " ")
	messageArray[0] = strings.ToLower(messageArray[0][1:])
	var commandList *commands.GenericCommand
	commandList = new(commands.GenericCommand)

	switch messageArray[0] {
	case "help":
		commandList.BuildHelpEmbed(session, message)
		break
	case "currency":
		commandList.GetCurrency(session, message, messageArray)
		break
	case "mute":
		commandList.MuteUser(session, message, messageArray, adminRole, mutedRole)
		break
	case "unmute":
		commandList.UnmuteUser(session, message, messageArray, adminRole, mutedRole)
		break
	case "mutes":
		commandList.ViewMutes(session, message, mutedRole)
		break
	}
}
