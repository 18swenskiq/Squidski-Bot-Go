package commands

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/translate"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/language"
)

func (c *GenericCommand) Translate(session *discordgo.Session, message *discordgo.MessageCreate, messageArray []string) {
	// Get client ready for translating
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	text := strings.Join(messageArray, " ")

	target, err := language.Parse("en")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Translation %v\n", translations[0].Text)

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00,
		Description: translations[0].Text,
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:       "Translating '" + text + "':",
	}
	session.ChannelMessageSendEmbed(message.ChannelID, embed)

}
