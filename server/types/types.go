package types

type Person struct {
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type Image struct {
	Url   string `json:"url" db:"url"`
	Title string `json:"title" db:"title"`
}

type Category struct {
	Term string `json:"term" db:"term"`
}

type Collection struct {
	User  string `json:"user" db:"user"`
	Title string `json:"title" db:"title"`
}

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

func NewCategory(term string) Category {
	return Category{Term: term}
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
