package models

type PlayerRankResponse struct {
	RankTier         int     `json:"rank_tier"`
	LeaderboardRank  int     `json:"leaderboard_rank"`
	ComputedMMR      float64 `json:"computed_mmr"`
	ComputedMMRTurbo float64 `json:"computed_mmr_turbo"`
	Aliases          []Alias `json:"aliases"`
	Profile          Profile `json:"profile"`
}

type Alias struct {
	PersonaName string `json:"personaname"`
	NameSince   string `json:"name_since"`
}

type Profile struct {
	AccountID      int    `json:"account_id"`
	PersonaName    string `json:"personaname"`
	Name           string `json:"name"`
	Plus           bool   `json:"plus"`
	Cheese         int    `json:"cheese"`
	SteamID        string `json:"steamid"`
	Avatar         string `json:"avatar"`
	AvatarMedium   string `json:"avatarmedium"`
	AvatarFull     string `json:"avatarfull"`
	ProfileURL     string `json:"profileurl"`
	LastLogin      string `json:"last_login"`
	LocCountryCode string `json:"loccountrycode"`
	IsContributor  bool   `json:"is_contributor"`
	IsSubscriber   bool   `json:"is_subscriber"`
}
