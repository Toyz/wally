package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/datasets"
	"strings"
)

func init() {
	Register(Command{
		Command:     "config",
		Desc:        "Control what I can send (Purity and Filter)",
		NSFW:        false,
		Func:        configCommand,
		HelpFunc: 	 configHelp,
		Permissions: discordgo.PermissionAdministrator,
	})
}

func configCommand(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate, args []string, config *datasets.Entity) error {
	if len(args) == 0 {
		embed := new(discordgo.MessageEmbed)
		embed.Title = "Wally Config"

		var keySet string
		if config.Guild.APIKey == "" {
			keySet = "False"
		} else {
			keySet = "True"
		}
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "API Key set",
			Value: keySet,
		})

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Filters",
			Value: config.Filter.Message(),
			Inline: true,
		})

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Purity",
			Value: config.Purity.Message(),
			Inline: true,
		})

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Enable/Disable Filters and Purities (Per Channel)",
			Value: "`w!config toggle {option}`\n\nOptions:\nGeneral, Anime, People, SFW, Sketchy, NSFW",
		})

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return err
	}

	if args[0] == "toggle" {
		if len(args) < 2 {
			return errors.New("Usage: w!config toggle {feature}")
		}

		switch strings.ToLower(args[1]) {
		case "general":
			config.Filter.General = !config.Filter.General
		case "anime":
			config.Filter.Anime = !config.Filter.Anime
		case "people":
			config.Filter.People = !config.Filter.People
		case "sfw":
			config.Purity.SFW = !config.Purity.SFW
		case "sketchy":
			if !c.NSFW {
				return errors.New("Can only be enabled in a NSFW channel")
			}
			config.Purity.Sketchy = !config.Purity.Sketchy
		case "nsfw":
			if !c.NSFW {
				return errors.New("Can only be enabled in a NSFW channel")
			}
			if config.Guild.APIKey == "" {
				return errors.New("API Key must be set in-order to set this... Get one at https://wallhaven.cc/settings/account")
			}
			config.Purity.NSFW = !config.Purity.NSFW
		}

		if err := storage.SetChannel(config.Channel); err != nil {
			return err
		}

		_, err := s.ChannelMessageSend(c.ID, "Updated config for current channel use `w!config` to view new config")
		return err
	}

	if args[0] == "set" {
		if len(args) < 3 {
			return errors.New("Usage: w!config set {key} {value}")
		}

		switch strings.ToLower(args[1]) {
		case "api_key":
			if args[2] == "none" {
				config.Guild.APIKey = ""
			} else {
				config.Guild.APIKey = args[2]
			}
		}

		if err := storage.SetGuild(config.Guild); err != nil {
			return err
		}

		_, err := s.ChannelMessageSend(c.ID, "Updated config for current channel use `w!config` to view new config")
		return err
	}

	return errors.New("Usage: w!config toggle {feature}")
}

func configHelp() *discordgo.MessageEmbed {
	embed := new(discordgo.MessageEmbed)
	embed.Title = "Wally Config Help"

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:  "API Key (Server Wide)",
		Value: "`w!config set api_key {key}`\n\nTo disabled `w!config set api_key none`",
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:  "Enable/Disable Filters and Purities (Per Channel)",
		Value: "`w!config toggle {option}`\n\nOptions:\nGeneral, Anime, People, SFW, Sketchy, NSFW",
	})

	return embed
}
