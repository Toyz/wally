package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/datasets"
	storage2 "github.com/toyz/wally/storage"
)

type CommandFunc func(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate, args []string, config *datasets.Entity) error
type HelpFunc func() *discordgo.MessageEmbed

type Command struct {
	Command     string
	Desc        string
	NSFW        bool
	Func        CommandFunc
	HelpFunc    HelpFunc
	Permissions int
}

var commands []Command
var storage storage2.Storage

func Register(cmd Command) {
	commands = append(commands, cmd)
}

func RegisterStorage(s storage2.Storage) {
	storage = s
}

func GetCommand(cmd string) (*Command, error) {
	for _, r := range commands {
		if r.Command == cmd {
			return &r, nil
		}
	}

	return nil, errors.New("unknown command")
}

func GetCommands() []Command {
	return commands
}
