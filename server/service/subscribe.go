package service

import (
	"github.com/mmcdole/gofeed"
	"github.com/pocketbase/pocketbase/daos"
)

type SubscribeService struct {
	dao *daos.Dao
}

func NewSubscribeService(dao *daos.Dao) *SubscribeService {
	return &SubscribeService{dao: dao}
}

func (s *SubscribeService) Subscribe(feed *gofeed.Feed) error {
	// return s.dao.Subscribe(feed)
	return nil
}
