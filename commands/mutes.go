package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	bolt "go.etcd.io/bbolt"
)

func (c *GenericCommand) ViewMutes(session *discordgo.Session, message *discordgo.MessageCreate, mutedRole string) {
	db, err := bolt.Open("../storage.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	nameValuesArray := []string{}
	valueValuesArray := []string{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MutedUsers"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			nameValuesArray = append(nameValuesArray, string(k))
			valueValuesArray = append(valueValuesArray, string(v))
		}
		return nil
	})
	stringToPrint := "\n"
	for i := range nameValuesArray {
		stringToPrint += "<@" + nameValuesArray[i] + "> is muted until " + valueValuesArray[i] + "\n"
	}
	stringToPrint += "\n"
	session.ChannelMessageSend(message.ChannelID, stringToPrint)
}
