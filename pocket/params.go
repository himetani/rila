package pocket

import "encoding/json"

type Params struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

func (p *Params) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}
