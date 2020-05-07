package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

type CommandFunc func(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate, args []string) error

type Command struct {
	Command     string
	Desc        string
	NSFW        bool
	Func        CommandFunc
	Permissions int
}

var commands []Command

func Register(cmd Command) {
	commands = append(commands, cmd)
}

func Get(cmd string) (*Command, error) {
	for _, r := range commands {
		if r.Command == cmd {
			return &r, nil
		}
	}

	return nil, errors.New("unknown command")
}
