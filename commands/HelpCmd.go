package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/datasets"
	"github.com/toyz/wally/tools/permissions"
)

func init() {
	Register(Command{
		Command:     "help",
		Desc:        "Show this",
		NSFW:        false,
		Func:        showHelp,
		Permissions: -1,
	})
}

func showHelp(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate, args []string, _ *datasets.Entity) error {
	if len(args) > 0 {
		command, err := GetCommand(args[0])
		if err != nil {
			return err
		}
		ok, err := permissions.MemberHasPermission(s, c.GuildID, m.Author.ID, command.Permissions)
		if err != nil {
			ok = false
		}
		if ok {
			if err != nil {
				return err
			}
			if command.HelpFunc != nil {
				_, err := s.ChannelMessageSendEmbed(m.ChannelID, command.HelpFunc())
				return err
			}
		}
	}

	embed := new(discordgo.MessageEmbed)
	embed.Title = "Help"

	for _, cmd := range GetCommands() {
		ok, err := permissions.MemberHasPermission(s, c.GuildID, m.Author.ID, cmd.Permissions)
		if err != nil {
			ok = false
		}
		if  ok {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   fmt.Sprintf("w!%s", cmd.Command),
				Value:  cmd.Desc,
				Inline: false,
			})
		}
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return err
}
