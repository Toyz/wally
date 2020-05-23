package commands

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/datasets"
	"github.com/toyz/wally/wallhaven"
)

func init() {
	Register(Command{
		Command:     "v",
		Desc:        "View a single image `!v 94x38z`",
		NSFW:        false,
		Func:        singleImage,
		Permissions: -1,
	})
}

func singleImage(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate, args []string, config *datasets.Entity, command Command) error {
	if len(args) == 0 {
		return errors.New("Invalid command usage `w!v <wallpaper id>` (example: `!v 94x38z`)")
	}

	image, err := wallhaven.SingleImage(config.Guild.APIKey, args[0])
	if err != nil {
		return err
	}

	if (image.Purity == "nsfw" || image.Purity == "sketchy") && !c.NSFW {
		return fmt.Errorf("Cannot display **%s** images in non-NSFW channel", image.Purity)
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, image.CreateEmbed())
	return err
}
