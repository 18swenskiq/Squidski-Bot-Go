package commands

import (
	"math/rand"
	"strconv"
	"time"

	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) GetSquidskiFact(session *discordgo.Session, message *discordgo.MessageCreate) {
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	_, keyValues := database.IterateOverKeysInBucketReturnBoth("SquidskiFacts")
	arrayIndex := rand.Intn(len(keyValues))

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00, // Green
		Description: keyValues[arrayIndex],
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/avatars/66318815247466496/31fd1ce7b18f33b5a5bd2f499c3cf477.png?size=64",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Squidski Fact #" + strconv.Itoa(arrayIndex),
	}
	session.ChannelMessageSendEmbed(message.ChannelID, embed)
}
