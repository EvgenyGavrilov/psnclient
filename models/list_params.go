package models

// ListParams параметры запроса получения списка игр от playstation store.
type ListParams struct {
	Size            int    `url:"size,omitempty"`
	Start           int    `url:"start"`
	Price           int    `url:"price,omitempty"`
	Sort            string `url:"sort,omitempty"`
	Direction       string `url:"direction,omitempty"`
	GameContentType string `url:"gameContentType,omitempty"`
	SubtitleLang    string `url:"subtitleLang,omitempty"`
	ReleaseDate     string `url:"releaseDate,omitempty"`
	Genre           string `url:"genre,omitempty"`
	TopCategory     string `url:"topCategory,omitempty"`
	VoiceLang       string `url:"voiceLang,omitempty"`
	GameType        string `url:"game_type,omitempty"`
	Relationship    string `url:"relationship,omitempty"`
	Platform        string `url:"platform,omitempty"`
	GameDemo        bool   `url:"gameDemo,omitempty"`
}
