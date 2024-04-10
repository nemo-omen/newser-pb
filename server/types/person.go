package types

import "encoding/json"

type Person struct {
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

func (p *Person) JSON() []byte {
	j, _ := json.MarshalIndent(p, "", "  ")
	return j
}

func (p *Person) String() string {
	return string(p.JSON())
}
