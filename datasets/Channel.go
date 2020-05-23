package datasets

type Channel struct {
	ChannelID string `json:"channel_id" bson:"channel_id"`
	Filter    Filter `json:"filter" bson:"filter"`
	Purity    Purity `json:"purity" bson:"purity"`
}

type Filter struct {
	General bool
	Anime   bool
	People  bool
}

type Purity struct {
	SFW     bool
	Sketchy bool
	NSFW    bool
}

// 111 (general/anime/people)
func (f Filter) String() string {
	filter := "000"
	if f.General {
		filter = replaceAtIndex(filter, '1', 0)
	}

	if f.Anime {
		filter = replaceAtIndex(filter, '1', 1)
	}

	if f.People {
		filter = replaceAtIndex(filter, '1', 2)
	}

	return filter
}

func (f Filter) Message() string {
	var purityPerms string
	if f.General {
		purityPerms += ":white_check_mark: General"
	} else {
		purityPerms += ":x: General"
	}
	purityPerms += "\n"

	if f.Anime {
		purityPerms += ":white_check_mark: Anime"
	} else {
		purityPerms += ":x: Anime"
	}
	purityPerms += "\n"

	if f.People {
		purityPerms += ":white_check_mark: People"
	} else {
		purityPerms += ":x: People"
	}

	return purityPerms
}


// 111 (SFW/Sketchy/NSFW)
func (f Purity) String() string {
	filter := "000"
	if f.SFW {
		filter = replaceAtIndex(filter, '1', 0)
	}

	if f.Sketchy {
		filter = replaceAtIndex(filter, '1', 1)
	}

	if f.NSFW {
		filter = replaceAtIndex(filter, '1', 2)
	}

	return filter
}

func (f Purity) Message() string {
	var purityPerms string
	if f.SFW {
		purityPerms += ":white_check_mark: SFW"
	} else {
		purityPerms += ":x: SFW"
	}
	purityPerms += "\n"

	if f.Sketchy {
		purityPerms += ":white_check_mark: Sketchy"
	} else {
		purityPerms += ":x: Sketchy"
	}
	purityPerms += "\n"

	if f.NSFW {
		purityPerms += ":white_check_mark: NSFW"
	} else {
		purityPerms += ":x: NSFW"
	}

	return purityPerms
}

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}
