package service

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"server/constant"

	"github.com/PuerkitoBio/goquery"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed"
)

var (
	ErrFailedGofeedParse = fmt.Errorf("gofeed failed to parse feed")
)

type SearchService struct {
	client *http.Client
}

func NewSearchService() *SearchService {
	client := &http.Client{}
	client.Timeout = 10 * time.Second

	return &SearchService{
		client: client,
	}
}

// IsValidSite tests whether a given url leads to
// a valid site by sending a HEAD request and
// returning whether the response StatusCode == 200
// If the request results in an error, the result is false.
func (s SearchService) IsValidSite(siteUrl string) bool {
	res, err := s.client.Head(siteUrl)
	if err != nil {
		return false
	}

	if res.StatusCode != http.StatusOK {
		return false
	}

	return true
}

// GetFeed attempts to retrieve a valid RSS/Atom/JSON feed
// from the given URL string.
//
// The URL string must include a scheme and must point to
// a resource that returns a valid feed. (ie: https:whateves.com/feed).
//
// GetFeed uses github.com/mmcdole/gofeed to make the request
// and parse the response body.
func (s SearchService) GetFeed(feedUrl string) (*gofeed.Feed, error) {
	goFeed := &gofeed.Feed{}
	fp := gofeed.NewParser()

	goFeed, err := fp.ParseURL(feedUrl)

	if err != nil {
		return nil, fmt.Errorf("GetFeed: fp.ParseURL: %w", ErrFailedGofeedParse)
	}

	// ensure feed siteurl is present
	if goFeed.Link == "" {
		// we already know we're getting a valid
		// url from the feedUrl, so we can use it
		// to set the feed's siteurl
		url, _ := url.Parse(feedUrl)
		if url != nil {
			scheme := url.Scheme
			host := url.Host
			goFeed.Link = scheme + "://" + host
		}
	}

	// ensure feed feedurl is present
	if goFeed.FeedLink == "" {
		goFeed.FeedLink = feedUrl
	}

	// ensure feed description is free of html tags
	goFeed.Description = strip.StripTags(goFeed.Description)

	// double-check for site favicon if feed.Image
	// is not present
	if goFeed.Image == nil {
		u, _ := url.Parse(feedUrl)
		if u != nil {
			if u.Path != "" {
				u.Path = "/"
			}

			link := u.String()
			src := s.GetFaviconSrc(link)
			// fmt.Println("src: ", src)

			if src != "" {
				srcUrl, _ := url.Parse(src)
				if srcUrl != nil {
					if srcUrl.Scheme == "" {
						srcUrl.Scheme = u.Scheme
					}

					if srcUrl.Host == "" {
						srcUrl.Host = u.Host
					}
					src = srcUrl.String()
				}

				goFeed.Image = &gofeed.Image{
					URL:   src,
					Title: goFeed.Title,
				}
			}
		}
	} else {
		if goFeed.Image.Title == "" {
			goFeed.Image.Title = goFeed.Title
		}
	}

	for _, item := range goFeed.Items {
		// strip and truncate item description
		item.Description = strip.StripTags(item.Description)
		if len(item.Description) > 256 {
			item.Description = item.Description[0:256] + "..."
		}
	}

	image := gofeed.Image{}
	if goFeed.Image != nil {
		image.URL = goFeed.Image.URL
		image.Title = goFeed.Image.Title
	}

	authors := []*gofeed.Person{}

	for _, author := range goFeed.Authors {
		authors = append(authors, &gofeed.Person{
			Name:  author.Name,
			Email: author.Email,
		})
	}

	items := []*gofeed.Item{}
	for _, goitem := range goFeed.Items {
		authors := []*gofeed.Person{}
		for _, author := range goitem.Authors {
			authors = append(authors, &gofeed.Person{
				Name:  author.Name,
				Email: author.Email,
			})
		}

		item := gofeed.Item{
			Title:           goitem.Title,
			Description:     goitem.Description,
			Content:         goitem.Content,
			Link:            goitem.Link,
			Links:           goitem.Links,
			Updated:         goitem.Updated,
			UpdatedParsed:   goitem.UpdatedParsed,
			Published:       goitem.Published,
			PublishedParsed: goitem.PublishedParsed,
			Authors:         authors,
			GUID:            goitem.GUID,
			Image:           &image,
			Categories:      goitem.Categories,
		}
		items = append(items, &item)
	}

	feed := gofeed.Feed{
		Title:           goFeed.Title,
		Description:     goFeed.Description,
		Link:            goFeed.Link,
		FeedLink:        goFeed.FeedLink,
		Links:           goFeed.Links,
		Updated:         goFeed.Updated,
		UpdatedParsed:   goFeed.UpdatedParsed,
		Published:       goFeed.Published,
		PublishedParsed: goFeed.PublishedParsed,
		Authors:         authors,
		Language:        goFeed.Language,
		Image:           &image,
		Categories:      goFeed.Categories,
		Items:           items,
		FeedType:        goFeed.FeedType,
	}

	return &feed, nil
}

func (s SearchService) GetFeeds(feedUrls []string) ([]*gofeed.Feed, error) {
	feeds := []*gofeed.Feed{}
	type Result struct {
		Res   *gofeed.Feed
		Error error
	}

	ch := make(chan Result, len(feedUrls))

	for _, link := range feedUrls {
		u := link
		go func() {
			res, err := s.GetFeed(u)
			if err != nil {
				ch <- Result{
					Res:   nil,
					Error: err,
				}
			} else {
				// substitute feed's defined
				// link for valid url we just got
				// the feed from. Why? Because
				// sometimes people forget to update
				// the url in their RSS feeds. If we've
				// found a valid feed at this URL, it
				// obviously works, so we may as well use it.
				res.FeedLink = u
				ch <- Result{
					Res:   res,
					Error: nil,
				}
			}
		}()
	}

	for i := 0; i < len(feedUrls); i++ {
		result := <-ch
		if result.Error == nil {
			feeds = append(feeds, result.Res)
		}
	}

	return feeds, nil
}

func (s SearchService) GetFaviconSrc(siteUrl string) string {
	src := ""
	res, err := s.client.Get(siteUrl)
	if err != nil {
		return src
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return src
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return src
	}

	document.Find("link").Each(func(i int, el *goquery.Selection) {
		rel, exists := el.Attr("rel")
		if !exists {
			return
		}
		if !strings.Contains(rel, "icon") {
			return
		}

		href, exists := el.Attr("href")

		if !exists {
			return
		}

		src = href
	})
	return src
}

// FindFeedLinks searches the document at a given URL for
// feed links.
// siteUrl should be a valid URL (ie: https://whateves.com)
func (s SearchService) FindFeedLinks(siteUrl string) ([]string, error) {
	fmt.Println("Attempting to find feed links at", siteUrl)
	links := []string{}

	validUrl, err := GetValidUrl(siteUrl)
	if err != nil {
		return links, fmt.Errorf("FindFeedLinks: %w", ErrInvalidUrl)
	}

	fmt.Println("validUrl: ", validUrl)
	fullURL, _ := url.Parse(validUrl)

	res, err := s.client.Get(fullURL.String())
	if err != nil {
		return links, fmt.Errorf("FindFeedLinks: error fetching site: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return links, fmt.Errorf("FindFeedLinks: bad status code: %d", res.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return links, fmt.Errorf("FindFeedLinks: error parsing document: %w", err)
	}

	document.Find("link").Each(func(i int, s *goquery.Selection) {
		rel, exists := s.Attr("rel")
		if !exists {
			return
		}

		if rel != "alternate" {
			return
		}

		linkType, exists := s.Attr("type")

		if !exists {
			return
		}

		if !slices.Contains(constant.ValidContentTypes, constant.ContentType(linkType)) {
			return
		}

		href, exists := s.Attr("href")
		if !exists {
			return
		}

		testUrl, err := url.Parse(href)
		if err != nil {
			return
		}

		if !testUrl.IsAbs() {
			testUrl.Scheme = fullURL.Scheme
			testUrl.Host = fullURL.Host
			links = append(links, testUrl.String())
		} else {
			links = append(links, href)
		}
	})

	return links, nil
}

// GuessFeedLinks attempts to guess the endpoint
// where an RSS/Atom/JSON feed lives given a valid
// URL (ie: https://siteurl.com).
// Note, it should be called only after after s.FindFeedLinks
// has failed
func (s SearchService) GuessFeedLinks(siteUrl string) ([]string, error) {
	confirmed := []string{}
	guesses := []string{}

	// for each common feed path
	// attempt to create a valid url
	// if valid, append to potentialGuesses
	for _, path := range constant.CommonFeedPaths {
		u, err := url.ParseRequestURI(siteUrl + path)
		if err != nil {
			return confirmed, fmt.Errorf("GuessFeedLinks: could not parse url %s: %w", siteUrl+path, err)
		}
		guesses = append(guesses, u.String())
	}

	// For each guess, append one of the common feed
	// extensions, as long as the guess does not already
	// end with "/" or one of the common extensions
	// for _, guess := range guesses {
	// 	for _, ext := range common.CommonFeedExtensions {
	// 		if !strings.HasSuffix(guess, "/") && !strings.HasSuffix(guess, ext) {
	// 			withExt := guess + ext
	// 			guesses = append(guesses, withExt)
	// 		}
	// 	}
	// }

	// Create len(guesses) channels
	// and make requests concurrently.
	// If we receive httpStatusOK &&
	// heades.Content-Type is in ValidContentTypes
	// we probably have a match
	type Result struct {
		Res   *http.Response
		Error error
	}
	ch := make(chan Result, len(guesses))

	for _, guess := range guesses {
		u := guess
		go func() {
			res, err := s.client.Get(u)
			if err != nil {
				ch <- Result{
					Res:   nil,
					Error: err,
				}
			}

			if res.StatusCode != http.StatusOK {
				ch <- Result{
					Res:   res,
					Error: fmt.Errorf("bad result"),
				}
			} else {
				ch <- Result{
					Res:   res,
					Error: nil,
				}
			}
		}()
	}

	for _, guess := range guesses {
		result := <-ch
		if result.Error == nil {
			res := result.Res

			if res.Body != nil {
				defer res.Body.Close()
				contentType := res.Header.Get("Content-Type")
				if slices.Contains(constant.ValidDocContentTypes, constant.DocContentType(contentType)) {
					fmt.Println("guess: ", guess)
					confirmed = append(confirmed, guess)
				}
			}
		}
	}
	return confirmed, nil
}
