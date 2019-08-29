package services

import (
	"fmt"
	"time"

	utilities "../utilities"
	"github.com/bwmarrin/discordgo"
)

type MuteService struct {
	session discordgo.Session
	message discordgo.MessageCreate
}

func (c *MuteService) MuteService(session *discordgo.Session, mutedRole string, serverID string) {

	var database *utilities.GeneralDB
	database = new(utilities.GeneralDB)

	for {
		time.Sleep(time.Second * 5)
		names, values := database.IterateOverKeysInBucketReturnBoth("MutedUsers")
		if len(names) == 0 {
			continue
		}
		for i, _ := range names {
			t, err := time.Parse("2006-01-02 15:04:05-07:00", values[i])
			if t.Before(time.Now()) {
				fmt.Println("Mute has expired for " + names[i])
				database.DeleteKey("MutedUsers", names[i])
				err = session.GuildMemberRoleRemove(serverID, names[i], mutedRole)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

	}
}
