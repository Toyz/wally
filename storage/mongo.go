package storage

import (
	"github.com/globalsign/mgo/bson"
	"github.com/toyz/wally/datasets"
	"github.com/toyz/wally/storage/mongo_helpers"
)

func init() {
	RegisterStorage("mongo", &monogStorage{})
}

type monogStorage struct {
	config StorageConfig
}

func (m *monogStorage) Setup(config StorageConfig) {
	mongo_helpers.New(mongo_helpers.MongoSettings{
		Host:     config.Addr,
		Username: config.Username,
		Password: config.Password,
		Database: config.Database,
	})
}

func (m *monogStorage) Close() error {
	return mongo_helpers.Close()
}

func (m *monogStorage) GetChannel(id string) (channel datasets.Channel, err error) {
	err = mongo_helpers.One("channel", bson.M{
		"channel_id": id,
	}, &channel)

	if channel.ChannelID == id {
		return
	}

	err = nil
	channel = channelDefault
	channel.ChannelID = id
	return
}

func (m *monogStorage) SetChannel(channel datasets.Channel) error {
	return mongo_helpers.Upsert("channel", bson.M{
		"channel_id": channel.ChannelID,
	}, channel)

}

func (m *monogStorage) GetGuild(id string) (guild datasets.Guild, err error) {
	err = mongo_helpers.One("guild", bson.M{
		"guild_id": id,
	}, &guild)

	if guild.GuildID == id {
		return
	}

	err = nil
	guild.GuildID = id
	return}

func (m *monogStorage) SetGuild(guild datasets.Guild) error {
	return mongo_helpers.Upsert("guild", bson.M{
		"guild_id": guild.GuildID,
	}, guild)
}
