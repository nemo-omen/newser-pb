package service

import (
	"fmt"
	"server/types"

	"github.com/pocketbase/pocketbase/daos"
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
		return s.createSubscription(userID, storedFeed.Id)
	}
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
	return "", nil
}
