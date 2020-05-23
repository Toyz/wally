package storage

import (
	"github.com/toyz/wally/datasets"
	"strings"
)

var channelDefault = datasets.Channel{
	Filter: datasets.Filter{
		General: true, Anime: true, People: true,
	},
	Purity: datasets.Purity{
		SFW: true,
	},
}

type StorageConfig struct {
	Username string
	Password string
	Addr     string
	Database string
}

type Storage interface {
	Setup(config StorageConfig)
	Close() error

	GetChannel(id string) (datasets.Channel, error)
	SetChannel(channel datasets.Channel) error

	GetGuild(id string) (datasets.Guild, error)
	SetGuild(id datasets.Guild) error
}

var types map[string]Storage

func RegisterStorage(name string, storage Storage) {
	if types == nil {
		types = make(map[string]Storage)
	}

	types[strings.ToLower(name)] = storage
}

func Get(name string) Storage {
	s, ok := types[strings.ToLower(name)]
	if !ok {
		panic("unknown storage container type")
	}

	return s
}
