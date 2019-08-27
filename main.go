package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Load the botkey
	botKeyBytes, err := ioutil.ReadFile("settings/botkey.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	// Convert the botkey from bytes to String
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&botKeyBytes))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	botKeyString := *(*string)(unsafe.Pointer(&sh))

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(botKeyString)
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return
	}

	// Register the messageCreate as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and being listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	// Wait here until CTRL-C or other term signal is recieved.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required but its good practice
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
