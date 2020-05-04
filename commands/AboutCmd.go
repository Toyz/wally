package commands

import "github.com/bwmarrin/discordgo"

func init() {
	Register("!about", aboutWally)
}

func aboutWally(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	embed := new(discordgo.MessageEmbed)
	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      "https://i.imgur.com/7EAX8Zi.gif",
		Width:    245,
		Height:   245,
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

	_, e := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return e
}