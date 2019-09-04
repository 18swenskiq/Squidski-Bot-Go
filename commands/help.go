package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// TODO: Create overrides to add a help for each command

func (c *GenericCommand) BuildHelpEmbed(session *discordgo.Session, message *discordgo.MessageCreate, adminRole string, messageArray []string) {

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00,
		Description: "Command List",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "**Add to Database (Admin Only):**",
				Value:  ">addtodb <Database Bucket> <KeyName> <KeyValue>\nAdds a value to the database",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Casino:**",
				Value:  ">c or >casino\nAccesses the casino module",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Currency:**",
				Value:  ">currency <amount> <currency code converting from>\nConverts currency",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Mute (Admin Only):**",
				Value:  ">mute <user> <time in minutes>\nSends a user to the void for a given amount of time",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Mutes:**",
				Value:  ">mutes\nView all the mutes in the server",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Pings:**",
				Value:  ">pings\nSubscribe or unsubscribe from the pings role",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Remove From Database (Admin only):**",
				Value:  ">removefromdb <Database Bucket> <KeyName>\nRemove a specific key from a database bucket",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Role Ping (Admin only):**",
				Value:  ">roleping\nPing the 'pings' role",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Seinfeldme:**",
				Value:  ">seinfeldme\nGet a random seinfeld quote",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Squidski Fact:**",
				Value:  ">squidskifact\nGet a random Squidski fact",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "**Unmute (Admin Only):**",
				Value:  ">unmute <user>\nUnmute a user",
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "Help Embed",
	}
	session.ChannelMessageSendEmbed(message.ChannelID, embed)
}
