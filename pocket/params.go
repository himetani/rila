package pocket

import "encoding/json"

type Params struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
	ContentType string `json:"contentType,omitempty"`
	Sort        string `json:"sort,omitempty"`
	Count       int    `json:"count,omitempty"`
}

func (p *Params) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}
