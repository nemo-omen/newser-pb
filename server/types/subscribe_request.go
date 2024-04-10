package types

import "encoding/json"

type SubscribeRequest struct {
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	FeedLink    string     `json:"feed_link"`
	SiteLink    string     `json:"site_link"`
	UpdatedAt   string     `json:"updated_at"`
	PublishedAt string     `json:"published_at"`
	Authors     []*Person  `json:"authors"`
	Language    string     `json:"language"`
	Copyright   string     `json:"copyright"`
	Image       *Image     `json:"image"`
	Categories  []string   `json:"categories"`
	FeedType    string     `json:"feed_type"`
	Articles    []*Article `json:"articles"`
}

func (s *SubscribeRequest) ToNewsfeed() Newsfeed {

	nAuthors := []*Person{}
	for _, a := range s.Authors {
		nAuthors = append(nAuthors, &Person{
			Name:  a.Name,
			Email: a.Email,
		})
	}

	categories := []*Category{}
	for _, c := range s.Categories {
		categories = append(categories, &Category{
			Term: c,
		})
	}

	articles := []*Article{}
	for _, a := range s.Articles {
		authors := []*Person{}
		for _, a := range a.Authors {
			authors = append(authors, &Person{
				Name:  a.Name,
				Email: a.Email,
			})
		}

		categories := []*Category{}
		for _, c := range a.Categories {
			categories = append(categories, &Category{
				Term: c.Term,
			})
		}

		articles = append(articles, &Article{
			Title:       a.Title,
			Description: a.Description,
			Content:     a.Content,
			Link:        a.Link,
			SiteLink:    a.SiteLink,
			UpdatedAt:   a.UpdatedAt,
			PublishedAt: a.PublishedAt,
			GUID:        a.GUID,
			Authors:     authors,
			Image: &Image{
				Url:   a.Image.Url,
				Title: a.Image.Title,
			},
			Categories: categories,
		})
	}

	return Newsfeed{
		Title:       s.Title,
		Description: s.Description,
		FeedLink:    s.FeedLink,
		SiteLink:    s.SiteLink,
		UpdatedAt:   s.UpdatedAt,
		PublishedAt: s.PublishedAt,
		Authors:     nAuthors,
		Language:    s.Language,
		Categories:  categories,
		Copyright:   s.Copyright,
		Image: &Image{
			Url:   s.Image.Url,
			Title: s.Image.Title,
		},
		FeedType: s.FeedType,
		Articles: articles,
	}
}

func (s *SubscribeRequest) JSON() []byte {
	j, _ := json.MarshalIndent(s, "", "  ")
	return j
}

func (s *SubscribeRequest) String() string {
	return string(s.JSON())
}
