package commands

import "github.com/bwmarrin/discordgo"

type GenericCommand struct {
	session discordgo.Session
	message discordgo.MessageCreate
}
