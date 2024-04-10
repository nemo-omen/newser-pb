package types

import "encoding/json"

type Image struct {
	Url   string `json:"url" db:"url"`
	Title string `json:"title" db:"title"`
}

func (i *Image) JSON() []byte {
	j, _ := json.MarshalIndent(i, "", "  ")
	return j
}

func (i *Image) String() string {
	return string(i.JSON())
}
