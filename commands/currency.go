package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/bwmarrin/discordgo"
	forex "github.com/g3kk0/go-forex"
)

func (c *GenericCommand) GetCurrency(session *discordgo.Session, message *discordgo.MessageCreate, messageArray []string) {
	// If not enough parameters are provided, yeet out of here
	if len(messageArray) != 3 {
		session.ChannelMessageSend(message.ChannelID, "Incorrect number of parameters! Format currency requests as "+string(message.Content[0])+"currency <amount> <currency code>")
		return
	}
	// Only continue if the user is trying to convert a float or int
	if !govalidator.IsFloat(messageArray[1]) {
		session.ChannelMessageSend(message.ChannelID, messageArray[1]+"is not a valid number.")
		return
	}
	// Check to make sure the currency that the user it trying to convert from is recognized by the API we're using
	allowedCurrencies := [33]string{"CAD", "HKD", "ISK", "PHP", "DKK", "HUF", "CZK", "AUD", "RON", "SEK", "IDR", "INR", "BRL", "RUB", "HRK", "JPY", "THB", "CHF", "EUR", "SGD", "PLN", "BGN", "TRY", "CNY", "NOK", "NZD", "ZAR", "USD", "MXN", "ILS", "GBP", "KRW", "MYR"}
	currencyContainsFlag := false
	for _, i := range allowedCurrencies {
		if strings.ToUpper(messageArray[2]) == i {
			currencyContainsFlag = true
		}
	}
	if currencyContainsFlag == false {
		session.ChannelMessageSend(message.ChannelID, messageArray[2]+"is not a recognized currency.")
		return
	}
	session.ChannelMessageSend(message.ChannelID, "Sending GET Request, please hold...")
	// Get the currency rates
	fc := forex.NewClient()
	params := map[string]string{"base": messageArray[2]}
	rates, err := fc.Latest(params)
	if err != nil {
		log.Println(err)
	}
	currencyInput, err := strconv.ParseFloat(messageArray[1], 64)
	if err != nil {
		log.Println(err)
	}
	sendmessage := "```py\nConverting " + messageArray[1] + " " + strings.ToUpper(messageArray[2]) + ":\n"
	sendmessage += "United States Dollar: " + fmt.Sprintf("%.2f", rates.Rates["USD"]*currencyInput) + "\n"
	sendmessage += "Australian Dollar: " + fmt.Sprintf("%.2f", rates.Rates["AUD"]*currencyInput) + "\n"
	sendmessage += "South African Rand: " + fmt.Sprintf("%.2f", rates.Rates["ZAR"]*currencyInput) + "\n"
	sendmessage += "Bulgarian Lev: " + fmt.Sprintf("%.2f", rates.Rates["BGN"]*currencyInput) + "\n"
	sendmessage += "Polish złoty: " + fmt.Sprintf("%.2f", rates.Rates["PLN"]*currencyInput) + "\n"
	sendmessage += "New Zealand Dollar: " + fmt.Sprintf("%.2f", rates.Rates["NZD"]*currencyInput) + "\n"
	sendmessage += "Great Britain Pound: " + fmt.Sprintf("%.2f", rates.Rates["GBP"]*currencyInput) + "\n"
	if strings.ToUpper(messageArray[2]) != "EUR" {
		sendmessage += "Euro: " + fmt.Sprintf("%.2f", rates.Rates["EUR"]*currencyInput) + "\n"
	} else {
		sendmessage += "Euro: " + messageArray[1] + "\n"
	}
	sendmessage += "```"
	session.ChannelMessageSend(message.ChannelID, sendmessage)
}
