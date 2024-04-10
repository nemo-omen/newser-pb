package types

import "encoding/json"

type Newsfeed struct {
	Title       string      `json:"title" db:"title"`
	Description string      `json:"description" db:"description"`
	FeedLink    string      `json:"feed_link" db:"feed_link"`
	SiteLink    string      `json:"site_link" db:"site_link"`
	UpdatedAt   string      `json:"updated_at" db:"updated_at"`
	PublishedAt string      `json:"published_at" db:"published_at"`
	Authors     []*Person   `json:"authors" db:"authors"`
	Language    string      `json:"language" db:"language"`
	Copyright   string      `json:"copyright" db:"copyright"`
	Image       *Image      `json:"image" db:"image"`
	Categories  []*Category `json:"categories" db:"categories"`
	FeedType    string      `json:"feed_type" db:"feed_type"`
	Articles    []*Article  `json:"articles" db:"articles"`
}

func (n *Newsfeed) JSON() []byte {
	j, _ := json.MarshalIndent(n, "", "  ")
	return j
}

func (n *Newsfeed) String() string {
	return string(n.JSON())
}
