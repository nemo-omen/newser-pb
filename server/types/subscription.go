package types

import "encoding/json"

type Subscription struct {
	UserID string `json:"user_id" db:"user_id"`
	FeedID string `json:"feed_id" db:"feed_id"`
}

type SubscriptionDTO struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	FeedID string `json:"feed_id"`
}

func (s *Subscription) JSON() []byte {
	j, _ := json.MarshalIndent(s, "", "  ")
	return j
}

func (s *Subscription) String() string {
	return string(s.JSON())
}
