package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/toyz/wally/wallhaven"
	"math/rand"
	"time"
)

func init() {
	Register(Command{
		Command:     "r",
		Desc:        "Random Image",
		NSFW:        false,
		Func:        randomImage(""),
		Permissions: discordgo.PermissionSendMessages | discordgo.PermissionAdministrator,
	})

	Register(Command{
		Command:     "rg",
		Desc:        "Random General Image",
		NSFW:        false,
		Func:        randomImage("100"),
		Permissions: discordgo.PermissionSendMessages | discordgo.PermissionAdministrator,
	})

	Register(Command{
		Command: "ra",
		Desc:    "Random Anime Image",
		NSFW:    false,
		Func:    randomImage("010"),
		Permissions: discordgo.PermissionSendMessages | discordgo.PermissionAdministrator,
	})

	Register(Command{
		Command: "rp",
		Desc:    "Random Person Image",
		NSFW:    false,
		Func:    randomImage("001"),
		Permissions: discordgo.PermissionSendMessages | discordgo.PermissionAdministrator,
	})
}

func randomImage(category string) CommandFunc {
	return func(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate, args []string) error {
		resolution := ""
		if len(args) > 0 {
			resolution = args[0]
		}

		papers, err := getWallPapers(category, "100", resolution)
		if err != nil {
			return err
		}

		paper := randomWallpaper(papers)
		paper, err = wallhaven.SingleImage("", paper.ID)
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
	papers, err := wallhaven.RandomImage("", category, purity, resolution)
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
