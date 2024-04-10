package types

import "encoding/json"

type Category struct {
	Term string `json:"term" db:"term"`
}

func NewCategory(term string) Category {
	return Category{Term: term}
}

func (c *Category) JSON() []byte {
	j, _ := json.MarshalIndent(c, "", "  ")
	return j
}

func (c *Category) String() string {
	return string(c.JSON())
}
