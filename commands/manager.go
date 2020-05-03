package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

type CommandFunc func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error

type Command struct {
	Command string
	Func    CommandFunc
}

var commands []Command

func Register(cmd string, action CommandFunc) {
	commands = append(commands, Command{
		Command: cmd,
		Func:    action,
	})
}

func Get(cmd string) (CommandFunc, error) {
	for _, r := range commands {
		if r.Command == cmd {
			return r.Func, nil
		}
	}

	return nil, errors.New("unknown command")
}