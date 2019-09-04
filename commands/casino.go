package commands

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	utilities "../utilities"
	"github.com/asaskevich/govalidator"
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

	// Parse our user's stats (here we're just pulling their stats from the DB and turning them into ints)
	userSquidCoins, _ := strconv.Atoi(strings.Split(keyValues[indexOfUserInLists], ",")[0])
	userTimesGambled, _ := strconv.Atoi(strings.Split(keyValues[indexOfUserInLists], ",")[1])
	userCoinsLost, _ := strconv.Atoi(strings.Split(keyValues[indexOfUserInLists], ",")[2])

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
					Value:  strconv.Itoa(userSquidCoins),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Times Gambled:",
					Value:  strconv.Itoa(userTimesGambled),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Coins Lost:",
					Value:  strconv.Itoa(userCoinsLost),
					Inline: true,
				},
			},
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "Squid Casino Stats",
		}
		session.ChannelMessageSendEmbed(message.ChannelID, embed)
		return
	case "roulette":
		// Must have 4 parameters to play roulette
		if len(messageArray) != 4 {
			session.ChannelMessageSend(message.ChannelID, "Incorrect number of parameters were provided. Type \""+string(message.Content[0])+"help c\" to view all possible casino commands")
			return
		}
		// Check to see if the user is betting on a valid number or phrase
		isBetPhrase := false
		if !govalidator.IsInt(messageArray[2]) {
			// Here we need to iterate over our valid bets to make sure its a valid word
			listOfValidBets := [5]string{"black", "red", "green", "evens", "odds"}
			for _, b := range listOfValidBets {
				if strings.ToLower(messageArray[2]) == b {
					isBetPhrase = true
					break
				}
			}
			if !isBetPhrase {
				session.ChannelMessageSend(message.ChannelID, messageArray[2]+" is not a valid thing you can bet on!")
				return
			}
		}
		// Check to make sure the amount they wagered is an integer amount of coins
		if !govalidator.IsInt(messageArray[3]) {
			session.ChannelMessageSend(message.ChannelID, messageArray[3]+" is not a number (or its just not an integer).")
			return
		}
		wageredAmount, _ := strconv.Atoi(messageArray[3])
		if wageredAmount < 1 {
			session.ChannelMessageSend(message.ChannelID, "Wagered amount cannot be a negative number or zero")
			return
		}
		if wageredAmount < userSquidCoins {
			session.ChannelMessageSend(message.ChannelID, "You cannot bet more coins than you have!")
			return
		}

		rouletteChoice := rand.Intn(36)
		chosenColor := "Black"
		redColors := [18]int{1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36}
		for _, i := range redColors {
			if rouletteChoice == i {
				chosenColor = "Red"
				break
			}
		}
		if rouletteChoice == 0 {
			chosenColor = "Green"
		}

		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0xB22222,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Number Landed On:",
					Value:  strconv.Itoa(rouletteChoice),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Color:",
					Value:  chosenColor,
					Inline: true,
				},
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Roulette Spin",
		}
		session.ChannelMessageSendEmbed(message.ChannelID, embed)
		chosenColor = strings.ToLower(chosenColor)
		userBet, _ := strconv.Atoi(messageArray[3])
		if !isBetPhrase {
			if messageArray[2] == strconv.Itoa(rouletteChoice) {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa((userSquidCoins-userBet)+userBet*36)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost))
				session.ChannelMessageSend(message.ChannelID, "Congrats! Your correct bet has netted you "+strconv.Itoa(userBet*36)+" Squidcoins for a total of "+strconv.Itoa((userSquidCoins-userBet)+userBet*36)+" coins!")
				return
			} else {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa(userSquidCoins-userBet)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost+userBet))
				session.ChannelMessageSend(message.ChannelID, "Your incorrect bet has lost you "+strconv.Itoa(userBet)+" Squidcoins for a total of "+strconv.Itoa(userSquidCoins-userBet)+" coins.")
				return
			}
		}
		switch strings.ToLower(messageArray[2]) {
		case "evens":
			if rouletteChoice%2 == 0 {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa((userSquidCoins-userBet)+userBet*2)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost))
				session.ChannelMessageSend(message.ChannelID, "Congrats! Your correct bet has netted you "+strconv.Itoa(userBet*2)+" Squidcoins for a total of "+strconv.Itoa((userSquidCoins-userBet)+userBet*2)+" coins!")
				return
			} else {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa(userSquidCoins-userBet)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost+userBet))
				session.ChannelMessageSend(message.ChannelID, "Your incorrect bet has lost you "+strconv.Itoa(userBet)+" Squidcoins for a total of "+strconv.Itoa(userSquidCoins-userBet)+" coins.")
				return
			}
		case "odds":
			if rouletteChoice%2 != 0 {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa((userSquidCoins-userBet)+userBet*2)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost))
				session.ChannelMessageSend(message.ChannelID, "Congrats! Your correct bet has netted you "+strconv.Itoa(userBet*2)+" Squidcoins for a total of "+strconv.Itoa((userSquidCoins-userBet)+userBet*2)+" coins!")
				return
			} else {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa(userSquidCoins-userBet)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost+userBet))
				session.ChannelMessageSend(message.ChannelID, "Your incorrect bet has lost you "+strconv.Itoa(userBet)+" Squidcoins for a total of "+strconv.Itoa(userSquidCoins-userBet)+" coins.")
				return
			}
		case "red":
			if chosenColor == "red" {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa((userSquidCoins-userBet)+userBet*2)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost))
				session.ChannelMessageSend(message.ChannelID, "Congrats! Your correct bet has netted you "+strconv.Itoa(userBet*2)+" Squidcoins for a total of "+strconv.Itoa((userSquidCoins-userBet)+userBet*2)+" coins!")
				return
			} else {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa(userSquidCoins-userBet)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost+userBet))
				session.ChannelMessageSend(message.ChannelID, "Your incorrect bet has lost you "+strconv.Itoa(userBet)+" Squidcoins for a total of "+strconv.Itoa(userSquidCoins-userBet)+" coins.")
				return
			}
		case "black":
			if chosenColor == "black" {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa((userSquidCoins-userBet)+userBet*2)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost))
				session.ChannelMessageSend(message.ChannelID, "Congrats! Your correct bet has netted you "+strconv.Itoa(userBet*2)+" Squidcoins for a total of "+strconv.Itoa((userSquidCoins-userBet)+userBet*2)+" coins!")
				return
			} else {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa(userSquidCoins-userBet)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost+userBet))
				session.ChannelMessageSend(message.ChannelID, "Your incorrect bet has lost you "+strconv.Itoa(userBet)+" Squidcoins for a total of "+strconv.Itoa(userSquidCoins-userBet)+" coins.")
				return
			}
		case "green":
			if chosenColor == "green" {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa((userSquidCoins-userBet)+userBet*2)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost))
				session.ChannelMessageSend(message.ChannelID, "Congrats! Your correct bet has netted you "+strconv.Itoa(userBet*2)+" Squidcoins for a total of "+strconv.Itoa((userSquidCoins-userBet)+userBet*2)+" coins!")
				return
			} else {
				database.WriteToDB("CasinoUsers", message.Author.ID, strconv.Itoa(userSquidCoins-userBet)+","+strconv.Itoa(userTimesGambled+1)+","+strconv.Itoa(userCoinsLost+userBet))
				session.ChannelMessageSend(message.ChannelID, "Your incorrect bet has lost you "+strconv.Itoa(userBet)+" Squidcoins for a total of "+strconv.Itoa(userSquidCoins-userBet)+" coins.")
				return
			}
		}
	}
}
