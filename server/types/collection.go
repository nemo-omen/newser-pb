package types

import "encoding/json"

type Collection struct {
	User  string `json:"user" db:"user"`
	Title string `json:"title" db:"title"`
}

func (c *Collection) JSON() []byte {
	j, _ := json.MarshalIndent(c, "", "  ")
	return j
}

func (c *Collection) String() string {
	return string(c.JSON())
}
