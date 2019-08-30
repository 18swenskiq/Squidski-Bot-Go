package commands

import (
	"math/rand"
	"strconv"
	"time"

	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) GetSeinfeldQuote(session *discordgo.Session, message *discordgo.MessageCreate) {
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	_, keyValues := database.IterateOverKeysInBucketReturnBoth("SeinfeldQuotes")
	arrayIndex := rand.Intn(len(keyValues))

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xe8fa1e,
		Description: keyValues[arrayIndex],
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       "Seinfeld Quote #" + strconv.Itoa(arrayIndex),
	}
	session.ChannelMessageSendEmbed(message.ChannelID, embed)
}
