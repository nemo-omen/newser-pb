package types

import "encoding/json"

type Article struct {
	Title       string      `json:"title" db:"title"`
	Description string      `json:"description" db:"description"`
	Content     string      `json:"content" db:"content"`
	Link        string      `json:"link" db:"link"`
	SiteLink    string      `json:"site_link" db:"site_link"`
	UpdatedAt   string      `json:"updated_at" db:"updated_at"`
	PublishedAt string      `json:"published_at" db:"published_at"`
	GUID        string      `json:"guid" db:"guid"`
	Authors     []*Person   `json:"authors" db:"authors"`
	Image       *Image      `json:"image" db:"image"`
	Categories  []*Category `json:"categories" db:"categories"`
}

func (a *Article) JSON() []byte {
	j, _ := json.MarshalIndent(a, "", "  ")
	return j
}

func (a *Article) String() string {
	return string(a.JSON())
}
