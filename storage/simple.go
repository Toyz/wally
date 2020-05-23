package storage

import (
	"errors"
	"fmt"
	"github.com/rapidloop/skv"
	"github.com/toyz/wally/datasets"
	"log"
)

func init() {
	RegisterStorage("simple", &simpleStorage{})
}

type simpleStorage struct {
	client *skv.KVStore
}

func (b *simpleStorage) Close() error {
	return b.client.Close()
}

func (b *simpleStorage) Setup(config StorageConfig) () {
	db, err := skv.Open(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	b.client = db
}

func (b *simpleStorage) GetChannel(id string) (channel datasets.Channel, err error) {
	err = b.client.Get(id, &channel)
	if err != nil {
		if !errors.Is(err, skv.ErrNotFound) {
			return
		}
	}

	if channel.ChannelID == id {
		return
	}

	// we need to reset error
	err = nil
	channel = channelDefault
	channel.ChannelID = id

	return
}

func (b *simpleStorage) SetChannel(channel datasets.Channel) error {
	return b.client.Put(channel.ChannelID, channel)
}

func (b *simpleStorage) GetGuild(id string) (guild datasets.Guild, err error) {
	key := fmt.Sprintf("guild_%s", id)
	err = b.client.Get(key, &guild)
	if err != nil {
		if !errors.Is(err, skv.ErrNotFound) {
			return
		}
	}

	if guild.GuildID == id {
		return
	}

	// we need to reset error
	err = nil
	guild.GuildID = id

	return
}

func (b *simpleStorage) SetGuild(guild datasets.Guild) error {
	key := fmt.Sprintf("guild_%s", guild.GuildID)

	return b.client.Put(key, guild)
}
