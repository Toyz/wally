package wallhaven

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Single struct {
	Data  Wallpaper `json:"data"`
	Error string    `json:"error"`
}

type Multi struct {
	Data []Wallpaper `json:"data"`
	Error string    `json:"error"`
}

type Avatar struct {
	Two00Px  string `json:"200px"`
	One28Px  string `json:"128px"`
	Three2Px string `json:"32px"`
	Two0Px   string `json:"20px"`
}

type Uploader struct {
	Username string `json:"username"`
	Group    string `json:"group"`
	Avatar   Avatar `json:"avatar"`
}

type Thumbs struct {
	Large    string `json:"large"`
	Original string `json:"original"`
	Small    string `json:"small"`
}

type Tags struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	CategoryID int    `json:"category_id"`
	Category   string `json:"category"`
	Purity     string `json:"purity"`
	CreatedAt  string `json:"created_at"`
}

type Wallpaper struct {
	ID         string   `json:"id"`
	URL        string   `json:"url"`
	ShortURL   string   `json:"short_url"`
	Uploader   Uploader `json:"uploader"`
	Views      int      `json:"views"`
	Favorites  int      `json:"favorites"`
	Source     string   `json:"source"`
	Purity     string   `json:"purity"`
	Category   string   `json:"category"`
	DimensionX int      `json:"dimension_x"`
	DimensionY int      `json:"dimension_y"`
	Resolution string   `json:"resolution"`
	Ratio      string   `json:"ratio"`
	FileSize   int      `json:"file_size"`
	FileType   string   `json:"file_type"`
	CreatedAt  string   `json:"created_at"`
	Colors     []string `json:"colors"`
	Path       string   `json:"path"`
	Thumbs     Thumbs   `json:"thumbs"`
	Tags       []Tags   `json:"tags"`
}

func (w Wallpaper) CreateEmbed() *discordgo.MessageEmbed {
	embed := new(discordgo.MessageEmbed)
	embed.Author = &discordgo.MessageEmbedAuthor{
		URL:          fmt.Sprintf("https://wallhaven.cc/user/%s", w.Uploader.Username),
		Name:         w.Uploader.Username,
	}

	//16748544
	if w.Purity == "sketchy" {
		embed.Color = 16748544
	}

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      w.Uploader.Avatar.Two00Px,
		Width:    200,
		Height:   200,
	}

	embed.Image = &discordgo.MessageEmbedImage{
		URL:      w.Path,
		Width:    w.DimensionX,
		Height:   w.DimensionY,
	}

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "Image Size",
		Value:  fmt.Sprintf("%dx%d", w.DimensionX, w.DimensionY),
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "View Image",
		Value:  w.ShortURL,
	})

	if w.Source != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "Image Source",
			Value:  w.Source,
		})
	}

	return embed
}