package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/datasets"
	"github.com/toyz/wally/wallhaven"
	"math/rand"
	"time"
)

func init() {
	Register(Command{
		Command:     "r",
		Desc:        "Random Image",
		NSFW:        false,
		Func:        randomImage(),
		Permissions: -1,
	})
}

func randomImage() CommandFunc {
	return func(s *discordgo.Session, _ *discordgo.Channel, m *discordgo.MessageCreate, args []string, config *datasets.Entity) error {
		resolution := ""
		if len(args) > 0 {
			resolution = args[0]
		}

		papers, err := getWallPapers(config.APIKey, config.Filter.String(), config.Purity.String(), resolution)
		if err != nil {
			return err
		}

		paper := randomWallpaper(papers)
		paper, err = wallhaven.SingleImage(config.Guild.APIKey, paper.ID)
		if err != nil {
			return err
		}

		_, err = s.ChannelMessageSendEmbed(m.ChannelID, paper.CreateEmbed())
		return err
	}
}

// 111 (general/anime/people)
func getWallPapers(apiKey, category, purity, resolution string) ([]wallhaven.Wallpaper, error) {
	if category == "" {

		category = "111"
	}
	papers, err := wallhaven.RandomImage(apiKey, category, purity, resolution)
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
