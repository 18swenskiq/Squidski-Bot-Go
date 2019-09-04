package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"unsafe"

	handlers "./handlers"
	services "./services"
	utilities "./utilities"

	"github.com/bwmarrin/discordgo"
)

type GeneralSettings struct {
	CallSymbol  string
	AdminRoleId string
	MutedRole   string
	ServerId    string
	PingsRole   string
	OwnerID     string
}

var globalCall = grabSettings().CallSymbol

func main() {
	// Load the botkey
	botKeyBytes, err := ioutil.ReadFile("settings/botkey.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Sucessfully opened botkey.txt")

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

	// Ensure our buckets in the DB exist
	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)
	database.EnsureBucketsExist("MutedUsers")
	database.EnsureBucketsExist("SquidskiFacts")
	database.EnsureBucketsExist("SeinfeldQuotes")
	database.EnsureBucketsExist("CasinoUsers")
	database.EnsureBucketsExist("KetalQuotes")
	database.EnsureBucketsExist("BruhMoments")

	// Print globalcall
	fmt.Println("The global call symbol is " + globalCall)

	// Start services
	var muteService *services.MuteService
	muteService = new(services.MuteService)
	go muteService.MuteService(dg, grabSettings().MutedRole, grabSettings().ServerId)

	// Register the messageCreate as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(guildMemberAdd)

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

// Get the general settings file
func grabSettings() GeneralSettings {
	settingsJson, err := ioutil.ReadFile("settings/general.json")
	if err != nil {
		fmt.Print(err)
	}
	var data GeneralSettings
	err = json.Unmarshal(settingsJson, &data)
	if err != nil {
		fmt.Println("error:", err)
	}
	return data
}

// Listen for users joining, then give them the pings role
func guildMemberAdd(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	err := s.GuildMemberRoleAdd(grabSettings().ServerId, event.User.ID, grabSettings().PingsRole)
	if err != nil {
		fmt.Println(err)
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required but its good practice
	if m.Author.ID == s.State.User.ID {
		return
	}
	if len(m.Content) < 1 {
		return
	}
	if m.Content[0] == byte(globalCall[0]) {
		var newCommand *handlers.CommandHandler
		newCommand = new(handlers.CommandHandler)
		data := grabSettings()
		newCommand.ExecuteCommand(s, m, data.AdminRoleId, data.MutedRole, data.PingsRole, data.OwnerID)
	}
}
