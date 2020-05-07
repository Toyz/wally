package commands

import "github.com/bwmarrin/discordgo"

func init() {
	Register(Command{
		Command: "!help",
		Desc:    "Show this",
		NSFW:    false,
		Func:    showHelp,
	})
}

func showHelp(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate, args []string) error {
	embed := new(discordgo.MessageEmbed)
	embed.Title = "Help"

	for _, cmd := range commands {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   cmd.Command,
			Value:  cmd.Desc,
			Inline: false,
		})
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return err
}
