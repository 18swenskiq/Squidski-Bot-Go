package commands

import (
	"fmt"

	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

func (c *GenericCommand) UnmuteUser(session *discordgo.Session, message *discordgo.MessageCreate, messageArray []string, adminRole string, mutedRole string) {

	if len(messageArray) != 2 {
		session.ChannelMessageSend(message.ChannelID, "Incorrect number of parameters! Format mute requests as "+string(message.Content[0])+"unmute <user>")
		return
	}
	if len(message.Mentions) != 1 {
		session.ChannelMessageSend(message.ChannelID, "No users were mentioned or too many were mentioned, so a mute cannot be applied")
		return
	}
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
		session.ChannelMessageSend(message.ChannelID, "You must be an administrator to use this command...")
		return
	}

	// Get info of the user we're muting to check if they already have the muted role
	unmutingUser, err := session.GuildMember(message.GuildID, message.Mentions[0].ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	checkIfAlreadyMutedFlag := false
	for _, i := range unmutingUser.Roles {
		if i == mutedRole {
			checkIfAlreadyMutedFlag = true
		}
	}
	if !checkIfAlreadyMutedFlag {
		session.ChannelMessageSend(message.ChannelID, "You cannot unmute a user that is not muted!")
		return
	}

	// Open the DB aand remove the mute
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	database.DeleteKey("MutedUsers", message.Mentions[0].ID)

	err = session.GuildMemberRoleRemove(message.GuildID, message.Mentions[0].ID, mutedRole)
	if err != nil {
		fmt.Println(err)
		return
	}
	session.ChannelMessageSend(message.ChannelID, "Unumted "+message.Mentions[0].Username)
}
