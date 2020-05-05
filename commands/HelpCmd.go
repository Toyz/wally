package commands

import "github.com/bwmarrin/discordgo"

func init() {
	Register("!help", "Show this", showHelp)
}

func showHelp(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	embed := new(discordgo.MessageEmbed)
	embed.Title = "Help"

	for _, cmd := range commands {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   cmd.Command,
			Value:  cmd.Desc,
			Inline: true,
		})
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return err
}
