package datasets

type Guild struct {
	GuildID string       `json:"guild_id" bson:"guild_id"`
	APIKey  string       `json:"api_key" bson:"api_key"`
	Options GuildOptions `json:"options" bson:"options"`
}

type GuildOptions struct {
	DisableAliases bool `bson:"disable_aliases"`
}
