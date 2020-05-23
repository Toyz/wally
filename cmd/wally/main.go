package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	commands2 "github.com/toyz/wally/commands"
	"github.com/toyz/wally/datasets"
	"github.com/toyz/wally/storage"
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
	Token        *string
	DatabaseType *string

	DatabaseAddr *string
	DatabaseDB   *string
	DatabaseUser *string
	DatabasePass *string

	Datastore storage.Storage
)

func init() {
	Token = kingpin.Flag("token", "Discord bot token").Short('t').Envar("DISCORD_BOT_TOKEN").Required().String()
	DatabaseType = kingpin.Flag("database_driver", "Database driver type").Short('d').Envar("DATABASE_DRIVER").Default("simple").String()
	DatabaseAddr = kingpin.Flag("database_addr", "Database address").Envar("DATABASE_ADDR").Default("").String()
	DatabaseDB = kingpin.Flag("database_name", "Database name").Envar("DATABASE_NAME").Default("./database.db").String()
	DatabaseUser = kingpin.Flag("database_user", "Database username").Envar("DATABASE_USER").Default("").String()
	DatabasePass = kingpin.Flag("database_pass", "Database password").Envar("DATABASE_PASS").Default("").String()

	kingpin.Parse()
}

func main() {
	Datastore = storage.Get(*DatabaseType)
	Datastore.Setup(storage.StorageConfig{
		Username: *DatabaseUser,
		Password: *DatabasePass,
		Addr:     *DatabaseAddr,
		Database: *DatabaseDB,
	})
	commands2.RegisterStorage(Datastore)
	defer Datastore.Close()

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

	base, command := m.Content[:2], m.Content[2:]
	if strings.EqualFold(base, "w!") {
		objs := strings.Split(command, " ")
		cmd, err := commands2.GetCommand(objs[0])
		if err != nil {
			log.Printf("Command was empty")
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

		channel, err := Datastore.GetChannel(m.ChannelID)
		if err != nil {
			log.Printf("failed to load channel config: %v", err)
			return
		}

		guilds, err := Datastore.GetGuild(m.GuildID)
		if err != nil {
			log.Printf("failed to load channel config: %v", err)
			return
		}

		if err := cmd.Func(s, c, m, objs[1:], &datasets.Entity{
			Channel: channel,
			Guild:   guilds,
		}); err != nil {
			log.Printf("failed to send messages: %v", err)
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
