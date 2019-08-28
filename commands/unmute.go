package commands

import (
	"fmt"
	"reflect"

	"github.com/bwmarrin/discordgo"
	bolt "go.etcd.io/bbolt"
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

	// Open the DB and check if the user is even muted
	db, err := bolt.Open("../storage.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MutedUsers"))
		c := b.Cursor()

		userAsByteArray := []byte(message.Mentions[0].ID)
		isUserMuted := []byte{0}
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			if reflect.DeepEqual(k, userAsByteArray) {
				isUserMuted = k
			}
		}
		if reflect.DeepEqual(isUserMuted, []byte{0}) {
			session.ChannelMessageSend(message.ChannelID, "This user's mute doesn't appear to be in the database, but I'll unmute them anyway")
			err = session.GuildMemberEdit(message.GuildID, message.Mentions[0].ID, []string{mutedRole})
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			b.Delete(userAsByteArray)
			err = session.GuildMemberRoleRemove(message.GuildID, message.Mentions[0].ID, mutedRole)
			if err != nil {
				fmt.Println(err)
				return err
			}
			session.ChannelMessageSend(message.ChannelID, "Sucessfully unmuted "+message.Mentions[0].Username)
		}

		return nil
	})
}
