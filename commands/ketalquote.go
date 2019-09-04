package commands

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) GetKetalQuote(session *discordgo.Session, message *discordgo.MessageCreate) {
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	_, keyValues := database.IterateOverKeysInBucketReturnBoth("KetalQuotes")
	arrayIndex := rand.Intn(len(keyValues))

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00, // Green
		Description: strings.Split(keyValues[arrayIndex], ";")[0],
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/avatars/342642223470608386/a76db7a1b97642687ebb7a1f015b07a5.png?size=512",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Ketal on " + strings.Split(keyValues[arrayIndex], ";")[1],
		},
		Title: "Ketal Quote #" + strconv.Itoa(arrayIndex),
	}
	session.ChannelMessageSendEmbed(message.ChannelID, embed)
}
