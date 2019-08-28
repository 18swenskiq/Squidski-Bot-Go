package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/bwmarrin/discordgo"
	bolt "go.etcd.io/bbolt"
)

func (c *GenericCommand) MuteUser(session *discordgo.Session, message *discordgo.MessageCreate, messageArray []string, adminRole string, mutedRole string) {
	// If not enough parameters are provided, yeet out of here
	if len(messageArray) != 3 {
		session.ChannelMessageSend(message.ChannelID, "Incorrect number of parameters! Format mute requests as "+string(message.Content[0])+"mute <user> <time>")
		return
	}

	// Make sure we're muting for a valid time
	if !govalidator.IsFloat(messageArray[2]) {
		session.ChannelMessageSend(message.ChannelID, messageArray[2]+"is not a valid number for mute times.")
		return
	}

	// Check to make sure a user is actually being referenced
	if len(message.Mentions) != 1 {
		session.ChannelMessageSend(message.ChannelID, "No users were mentioned or too many were mentioned, so a mute cannot be applied")
		return
	}

	// Get Info of user calling the mute so we can make sure they are an admin
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
	mutingUser, err := session.GuildMember(message.GuildID, message.Mentions[0].ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	checkIfAlreadyMutedFlag := false
	for _, i := range mutingUser.Roles {
		if i == mutedRole {
			checkIfAlreadyMutedFlag = true
		}
	}
	if checkIfAlreadyMutedFlag {
		session.ChannelMessageSend(message.ChannelID, "You cannot mute a user that is already muted!")
		return
	}

	// We need this here to set the mute duration
	timeMutedFloat, err := strconv.ParseFloat(messageArray[2], 32)
	timeMutedUntil := time.Now()
	timeMutedUntil = timeMutedUntil.Add(time.Minute * time.Duration(timeMutedFloat))

	// Let's open up the DB so we can store the mute
	db, err := bolt.Open("../storage.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("MutedUsers"))
		b := tx.Bucket([]byte("MutedUsers"))
		err := b.Put([]byte(message.Mentions[0].ID), []byte(timeMutedUntil.String()))
		return err
	})
	//db.View(func(tx *bolt.Tx) error {
	//b := tx.Bucket([]byte("MutedUsers"))
	//v := b.Get([]byte(message.Mentions[0].ID))
	//fmt.Printf("The user is muted until: %d\n", v)
	//return nil
	//})
	err = session.GuildMemberEdit(message.GuildID, message.Mentions[0].ID, []string{mutedRole})
	if err != nil {
		fmt.Println(err)
		return
	}
	session.ChannelMessageSend(message.ChannelID, "Successfully muted "+message.Mentions[0].Username+" for "+messageArray[2]+" minutes.")
}
