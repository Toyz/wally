package commands

import (
	"errors"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/datasets"
	"github.com/toyz/wally/wallhaven"
)

func init() {
	Register(Command{
		Command:     "r",
		Desc:        "Random image from channel defined filters",
		NSFW:        false,
		Func:        randomImage(),
		Permissions: -1,
	})

	Register(Command{
		Command:     "rg",
		Desc:        "Random general image",
		NSFW:        false,
		Alias:       true,
		Func:        randomImage(),
		Permissions: -1,
	})

	Register(Command{
		Command:     "ra",
		Desc:        "Random image from anime",
		NSFW:        false,
		Alias:       true,
		Func:        randomImage(),
		Permissions: -1,
	})

	Register(Command{
		Command:     "rp",
		Desc:        "Random image of a person",
		NSFW:        false,
		Alias:       true,
		Func:        randomImage(),
		Permissions: -1,
	})
}

func randomImage() CommandFunc {
	return func(s *discordgo.Session, _ *discordgo.Channel, m *discordgo.MessageCreate, args []string, config *datasets.Entity, command Command) error {
		filter := config.Filter.String()
		lastRune := command.Command[len(command.Command)-1]
		switch lastRune {
		case 'g':
			if config.Channel.Filter.General {
				filter = "100"
			} else {
				return errors.New("Command alias is disabled by filter rules")
			}
		case 'a':
			if config.Channel.Filter.Anime {
				filter = "010"
			} else {
				return errors.New("Command alias is disabled by filter rules")
			}
		case 'p':
			if config.Channel.Filter.People {
				filter = "001"
			} else {
				return errors.New("Command alias is disabled by filter rules")
			}
		}

		resolution := ""
		if len(args) > 0 {
			resolution = args[0]
		}

		papers, err := getWallPapers(filter, config.Purity.String(), resolution)
		if err != nil {
			return err
		}

		paper := randomWallpaper(papers)
		paper, err = wallhaven.SingleImage(paper.ID)
		if err != nil {
			return err
		}

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, paper.CreateEmbed())
		return err
	}
}

// 111 (general/anime/people)
func getWallPapers(category, purity, resolution string) ([]wallhaven.Wallpaper, error) {
	if category == "" {

		category = "111"
	}
	papers, err := wallhaven.RandomImage(category, purity, resolution)
	if err != nil {
		return nil, err
	}
	if len(papers) == 0 {
		return nil, errors.New("No Wallpapers have been found")
	}

	return papers, nil
}

func randomWallpaper(papers []wallhaven.Wallpaper) wallhaven.Wallpaper {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator

	return papers[r.Intn(len(papers))]
}
