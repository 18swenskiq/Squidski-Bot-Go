package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) GetTimeZones(session *discordgo.Session, message *discordgo.MessageCreate) {
	//Locations
	newZealand, _ := time.LoadLocation("Pacific/Auckland")
	easternAustralia, _ := time.LoadLocation("Australia/Sydney")
	bulgaria, _ := time.LoadLocation("Europe/Sofia")
	generalEuropeTime, _ := time.LoadLocation("Europe/Prague")
	easternTime, _ := time.LoadLocation("America/Fort_Wayne")
	centralTime, _ := time.LoadLocation("America/Chicago")
	pacificTime, _ := time.LoadLocation("America/Los_Angeles")

	timeObject := time.Now()
	customString := "```py\n"
	customString += "NZST: " + timeObject.In(newZealand).Format("15:04:05") + " [Peaches]\n"
	customString += "AEST: " + timeObject.In(easternAustralia).Format("15:04:05") + " [FluffyXVI, CommonCrayon, JackT]\n"
	customString += "EEST: " + timeObject.In(bulgaria).Format("15:04:05") + " [KlixX]\n"
	customString += "CEST: " + timeObject.In(generalEuropeTime).Format("15:04:05") + " [Squinky, Slimek, Florian]\n"
	customString += "EST:  " + timeObject.In(easternTime).Format("15:04:05") + " [Squidski, Seth, Skratchpost, CSGO John Madden, Strangest Stranger, jamdoggie]\n"
	customString += "CST:  " + timeObject.In(centralTime).Format("15:04:05") + " [TopHATTwaffle, Kale, SUMOSUPREME]\n"
	customString += "PST:  " + timeObject.In(pacificTime).Format("15:04:05") + " [Ch(i)ef, Practical Problems]\n"
	customString += "```"
	session.ChannelMessageSend(message.ChannelID, customString)

}
