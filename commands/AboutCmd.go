package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/datasets"
)

func init() {
	Register(Command{
		Command:     "about",
		Desc:        "About me",
		NSFW:        false,
		Func:        aboutWally,
		Permissions: -1,
	})
}

func aboutWally(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate, _ []string, _ *datasets.Entity, command Command) error {
	embed := new(discordgo.MessageEmbed)
	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:    "https://i.imgur.com/7EAX8Zi.gif",
		Width:  245,
		Height: 245,
	}
	embed.Title = "I am Wally"
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "What I do",
		Value:  "I make it so you can get fancy wallpapers from **Wallhaven.cc** right in discord!",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "Source code",
		Value:  "https://github.com/Toyz/wally",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "Created by,",
		Value:  "Helba (https://netslum.dev)",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:  "Invite me",
		Value: "https://discordapp.com/api/oauth2/authorize?client_id=706563357116727397&permissions=83968&scope=bot",
	})
	_, e := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return e
}
