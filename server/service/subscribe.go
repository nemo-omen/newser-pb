package service

import (
	"fmt"
	"server/types"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type SubscribeService struct {
	dao *daos.Dao
}

func NewSubscribeService(dao *daos.Dao) *SubscribeService {
	return &SubscribeService{dao: dao}
}

func (s *SubscribeService) Subscribe(subscribeReq *types.SubscribeRequest) error {
	userID := subscribeReq.UserID
	requestFeed := subscribeReq.ToNewsfeed()
	storedFeed, err := s.dao.FindFirstRecordByData("newsfeeds", "feed_link", requestFeed.FeedLink)
	if err == nil {
		// feed found, create & return subscription
		return s.createSubscription(userID, storedFeed.Id)
	}

	// feed not found, save newsfeed, then create subscription
	feedID, err := s.saveNewsfeed(requestFeed)

	if err != nil {
		return fmt.Errorf("failed to save newsfeed: %w", err)
	}
	return s.createSubscription(userID, feedID)
}

func (s *SubscribeService) createSubscription(userID, feedID string) error {
	subscription := types.Subscription{
		UserId: userID,
		FeedId: feedID,
	}
	_ = subscription
	// _, err := s.dao.CreateRecord("subscriptions", subscription)
	return nil
}

func (s *SubscribeService) saveNewsfeed(feed types.Newsfeed) (feedID string, err error) {
	newsfeedCollection, err := s.dao.FindCollectionByNameOrId("newsfeeds")
	if err != nil {
		return "", fmt.Errorf("failed to find newsfeeds collection: %w", err)
	}

	imageCollection, err := s.dao.FindCollectionByNameOrId("images")
	if err != nil {
		return "", fmt.Errorf("failed to find images collection: %w", err)
	}

	peopleCollection, err := s.dao.FindCollectionByNameOrId("people")
	if err != nil {
		return "", fmt.Errorf("failed to find people collection: %w", err)
	}

	categoryCollection, err := s.dao.FindCollectionByNameOrId("categories")
	if err != nil {
		return "", fmt.Errorf("failed to find categories collection: %w", err)
	}

	articleCollection, err := s.dao.FindCollectionByNameOrId("articles")
	if err != nil {
		return "", fmt.Errorf("failed to find articles collection: %w", err)
	}

	newsfeedRecord := models.NewRecord(newsfeedCollection)

	err = s.dao.RunInTransaction(func(txDao *daos.Dao) error {
		newsfeedRecord.Set("title", feed.Title)
		newsfeedRecord.Set("description", feed.Description)
		newsfeedRecord.Set("feed_link", feed.FeedLink)
		newsfeedRecord.Set("site_link", feed.SiteLink)
		newsfeedRecord.Set("updated_at", feed.UpdatedAt)
		newsfeedRecord.Set("published_at", feed.PublishedAt)
		newsfeedRecord.Set("language", feed.Language)
		newsfeedRecord.Set("copyright", feed.Copyright)
		newsfeedRecord.Set("feed_type", feed.FeedType)

		// save image
		imageID := ""
		if feed.Image != nil {
			imageID, err = saveImage(feed.Image, imageCollection, txDao)
			if err != nil {
				return fmt.Errorf("failed to save image: %w", err)
			}
		}
		newsfeedRecord.Set("image", imageID)

		// save authors
		authorIDs := []string{}
		for _, author := range feed.Authors {
			authorID, err := savePerson(author, peopleCollection, txDao)
			if err != nil {
				return fmt.Errorf("failed to save author: %w", err)
			}
			authorIDs = append(authorIDs, authorID)
		}
		newsfeedRecord.Set("authors", authorIDs)

		// save categories
		categoryIDs := []string{}
		for _, category := range feed.Categories {
			categoryID, err := saveCategory(category, categoryCollection, txDao)
			if err != nil {
				return fmt.Errorf("failed to save category: %w", err)
			}
			categoryIDs = append(categoryIDs, categoryID)
		}
		newsfeedRecord.Set("categories", categoryIDs)

		// save articles
		articleIDs := []string{}
		for _, article := range feed.Articles {
			articleID, err := saveArticle(article, articleCollection, txDao)
			if err != nil {
				return fmt.Errorf("failed to save article: %w", err)
			}
			articleIDs = append(articleIDs, articleID)
		}
		newsfeedRecord.Set("articles", articleIDs)

		if err := txDao.SaveRecord(newsfeedRecord); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to save newsfeed: %w", err)
	}

	return newsfeedRecord.Id, nil
}

func saveImage(image *types.Image, collection *models.Collection, dao *daos.Dao) (string, error) {
	imgRecord := models.NewRecord(collection)
	imgRecord.Set("url", image.Url)
	imgRecord.Set("title", image.Title)
	if err := dao.SaveRecord(imgRecord); err != nil {
		return "", err
	}
	return imgRecord.Id, nil
}

func savePerson(person *types.Person, collection *models.Collection, dao *daos.Dao) (string, error) {
	personRecord := models.NewRecord(collection)
	personRecord.Set("name", person.Name)
	personRecord.Set("email", person.Email)
	if err := dao.SaveRecord(personRecord); err != nil {
		return "", err
	}
	return personRecord.Id, nil
}

func saveArticle(article *types.Article, collection *models.Collection, dao *daos.Dao) (string, error) {
	imgCollection, err := dao.FindCollectionByNameOrId("images")
	if err != nil {
		return "", err
	}

	personCollection, err := dao.FindCollectionByNameOrId("people")
	if err != nil {
		return "", err
	}

	categoryCollection, err := dao.FindCollectionByNameOrId("categories")
	if err != nil {
		return "", err
	}

	articleRecord := models.NewRecord(collection)
	articleRecord.Set("title", article.Title)
	articleRecord.Set("description", article.Description)
	articleRecord.Set("content", article.Content)
	articleRecord.Set("link", article.Link)
	articleRecord.Set("site_link", article.SiteLink)
	articleRecord.Set("updated_at", article.UpdatedAt)
	articleRecord.Set("published_at", article.PublishedAt)
	articleRecord.Set("guid", article.GUID)

	// save image
	imageID := ""
	if article.Image != nil {
		imageID, err = saveImage(article.Image, imgCollection, dao)
		if err != nil {
			return "", err
		}
	}
	articleRecord.Set("image", imageID)

	// save authors
	authorIDs := []string{}
	for _, author := range article.Authors {
		authorID, err := savePerson(author, personCollection, dao)
		if err != nil {
			return "", err
		}
		authorIDs = append(authorIDs, authorID)
	}
	articleRecord.Set("authors", authorIDs)

	// save categories
	categoryIDs := []string{}
	for _, category := range article.Categories {
		categoryID, err := saveCategory(category, categoryCollection, dao)
		if err != nil {
			return "", err
		}
		categoryIDs = append(categoryIDs, categoryID)
	}
	articleRecord.Set("categories", categoryIDs)

	if err := dao.SaveRecord(articleRecord); err != nil {
		return "", err
	}
	return articleRecord.Id, nil
}

func saveCategory(category *types.Category, collection *models.Collection, dao *daos.Dao) (string, error) {
	categoryRecord := models.NewRecord(collection)
	categoryRecord.Set("name", category.Term)
	if err := dao.SaveRecord(categoryRecord); err != nil {
		return "", err
	}
	return categoryRecord.Id, nil
}
