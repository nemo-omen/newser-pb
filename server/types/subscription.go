package types

import "encoding/json"

type Subscription struct {
	UserId string `json:"user_id" db:"user_id"`
	FeedId string `json:"feed_id" db:"feed_id"`
}

func (s *Subscription) JSON() []byte {
	j, _ := json.MarshalIndent(s, "", "  ")
	return j
}

func (s *Subscription) String() string {
	return string(s.JSON())
}
