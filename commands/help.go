package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// TODO: Create overrides to add a help for each command

func (c *GenericCommand) BuildHelpEmbed(session *discordgo.Session, message *discordgo.MessageCreate, adminRole string) {
	adminUser, err := session.GuildMember(message.GuildID, message.Author.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	checkForAdminFlag := false
	for _, i := range adminUser.Roles {
		if i == adminRole {
			checkForAdminFlag = true
		}
	}
	if !checkForAdminFlag {
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x00ff00, // Green
			Description: "I am a bot made by Squidski#9545. I can do multiple things and I am still in development",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "I am a normal",
					Value:  "I am a value",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "I am a second field",
					Value:  "I am a value",
					Inline: true,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
			},
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "I am an Embed",
		}
		session.ChannelMessageSendEmbed(message.ChannelID, embed)
	} else {
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x00ff00, // Green
			Description: "I am a bot made by Squidski#9545. I can do multiple things and I am still in development",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "I am an admin",
					Value:  "I am a value",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "I am a second field",
					Value:  "I am a value",
					Inline: true,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
			},
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "I am an Embed",
		}
		session.ChannelMessageSendEmbed(message.ChannelID, embed)
	}
}
