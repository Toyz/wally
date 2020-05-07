package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/commands"
	"github.com/toyz/wally/tools/permissions"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	Token *string
	APIKey *string
)

func init() {
	Token = kingpin.Flag("token", "Discord bot token").Short('t').Envar("DISCORD_BOT_TOKEN").Required().String()
	APIKey = kingpin.Flag("api_key", "Wallhaven API Key").Short('a').Envar("WALLHAVEN_API_KEY").String()
	kingpin.Parse()
}

func main() {
	dg, err := discordgo.New(fmt.Sprintf("Bot %s", *Token))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if len(m.Content) < 3 {
		return
	}
	
	base, command := m.Content[0:2], m.Content[2:]
	if strings.EqualFold(base, "w!") {
		objs := strings.Split(command, " ")
		cmd, err := commands.Get(objs[0])
		if err != nil {
			return
		}

		c, err := s.Channel(m.ChannelID)
		if err != nil {
			log.Printf("failed to get channel: %v", err)
			return
		}

		if ok, err := permissions.MemberHasPermission(s, c.GuildID, m.Author.ID, cmd.Permissions); err != nil {
			log.Printf("failed to get user: %v", err)
			return
		} else {
			if !ok {
				_, _ = s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command")
				return
			}
		}

		if err := cmd.Func(s, c, m, objs[1:]); err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
		}
	}
}

func ready(s *discordgo.Session, _ *discordgo.Ready) {
	log.Print("Bot is now running. Press CTRL-C to exit.")

	s.UpdateStatus(0, "w!help || Eveeeee!")
	go func() {
		for {
			s.UpdateStatus(0, "w!help || Eveeeee!")
			time.Sleep(1 * time.Hour)
		}
	}()
}

