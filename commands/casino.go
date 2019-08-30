package commands

import (
	"strings"
	"time"

	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) UseCasino(session *discordgo.Session, message *discordgo.MessageCreate, messageArray []string) {
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)

	// Check all userdata in the casino, then open an account for someone if they haven't used the casino before
	nameValues, keyValues := database.IterateOverKeysInBucketReturnBoth("CasinoUsers")
	hasUsedCasinoBefore := false
	indexOfUserInLists := 0
	for i, user := range nameValues {
		if message.Author.ID == user {
			hasUsedCasinoBefore = true
			indexOfUserInLists = i
			break
		}
	}
	if !hasUsedCasinoBefore {
		// We don't need to extract all of our values from the database for this query again, since we can just append the new stuff in the DB to the slice we got earlier
		nameValues = append(nameValues, message.Author.ID)
		keyValues = append(keyValues, "1000,0,0")
		session.ChannelMessageSend(message.ChannelID, "I see you've never used the casino before. Welcome. I will give you 1000 Squidcoins to start off with.")
		database.WriteToDB("CasinoUsers", message.Author.ID, "1000,0,0")
	}

	// If not enough arguments were provided, we're not even gonna go into the casino command parsr
	if len(messageArray) < 2 {
		session.ChannelMessageSend(message.ChannelID, "No arguments to the casino module were provided. Type \""+string(message.Content[0])+"help c\" to view all possible casino commands")
		return
	} else {
		messageArray[1] = strings.ToLower(messageArray[1])
	}

	// Parse our user's stats
	userSquidCoins := strings.Split(keyValues[indexOfUserInLists], ",")[0]
	userTimesGambled := strings.Split(keyValues[indexOfUserInLists], ",")[1]
	userCoinsLost := strings.Split(keyValues[indexOfUserInLists], ",")[2]

	// Now we parse the user input to see what they wanted from the casino module
	switch messageArray[1] {
	// View Stats
	case "stats":
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xb53822,
			Description: "Casino stats for " + message.Author.Username,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Squidcoins:",
					Value:  userSquidCoins,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Times Gambled:",
					Value:  userTimesGambled,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Coins Lost:",
					Value:  userCoinsLost,
					Inline: true,
				},
			},
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "Squid Casino Stats",
		}
		session.ChannelMessageSendEmbed(message.ChannelID, embed)
		return
	case "roulette":

	}
}
