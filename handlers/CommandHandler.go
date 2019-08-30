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

func (c *CommandHandler) ExecuteCommand(session *discordgo.Session, message *discordgo.MessageCreate, adminRole string, mutedRole string, pingsRole string, ownerID string) {
	messageArray := strings.Split(message.Content, " ")
	messageArray[0] = strings.ToLower(messageArray[0][1:])
	var commandList *commands.GenericCommand
	commandList = new(commands.GenericCommand)

	switch messageArray[0] {
	case "help":
		commandList.BuildHelpEmbed(session, message, adminRole)
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
		commandList.ViewMutes(session, message)
		break
	case "pings":
		commandList.ChangePingState(session, message, pingsRole)
		break
	case "roleping":
		commandList.PingRole(session, message, adminRole, pingsRole)
		break
	case "addtodb":
		commandList.AddToDB(session, message, messageArray, ownerID)
		break
	case "squidskifact":
		commandList.GetSquidskiFact(session, message)
		break
	case "removefromdb":
		commandList.RemoveFromDB(session, message, messageArray, ownerID)
		break
	case "seinfeldme":
		commandList.GetSeinfeldQuote(session, message)
		break
	case "casino":
		commandList.UseCasino(session, message, messageArray)
		break
	case "c":
		commandList.UseCasino(session, message, messageArray)
		break
	}
}
