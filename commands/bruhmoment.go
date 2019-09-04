package commands

import (
	"math/rand"
	"strconv"

	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) BruhMoment(session *discordgo.Session, message *discordgo.MessageCreate) {
	randomValue := rand.Float32()
	if randomValue > .3 {
		session.ChannelMessageSend(message.ChannelID, "I have determined that this is not a bruh moment")
		return
	}
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	currentBruhMoments, _ := strconv.Atoi(database.ReadKey("BruhMoments", "NumberOfMoments"))
	database.WriteToDB("BruhMoments", "NumberOfMoments", strconv.Itoa(currentBruhMoments+1))
	session.ChannelMessageSend(message.ChannelID, "**THIS IS A CERTIFIED BRUH MOMENT**\nThis server has had "+strconv.Itoa(currentBruhMoments+1)+" bruh moments.")
	return
}
