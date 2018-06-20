package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type rescode struct {
	Code string `json:"code"`
}

type resAuth struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

type pocket struct {
	c *http.Client
}

func NewPocket() *pocket {
	return &pocket{
		c: &http.Client{},
	}
}

func (p *pocket) newReq(path, jsonstr string) (*http.Request, error) {
	req, err := http.NewRequest("POST", "https://getpocket.com/"+path, bytes.NewBuffer([]byte(jsonstr)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Accept", "application/json")

	return req, nil
}

func (p *pocket) GetRequestCode(key, redirect string) (string, error) {
	jsonstr := "{\"consumer_key\": \"" + key + "\", \"redirect_uri\": \"" + redirect + "\"}"
	log.Printf("Info: %s", jsonstr)

	req, err := p.newReq("v3/oauth/request", jsonstr)
	if err != nil {
		return "", err
	}

	resp, err := p.c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	//r = io.TeeReader(r, os.Stderr)
	fmt.Println("")

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status code was not 200. got=%d", resp.StatusCode)
	}

	var rc rescode
	if err := json.NewDecoder(r).Decode(&rc); err != nil {
		return "", err
	}
	return rc.Code, nil
}

func (p *pocket) GetAccessToken(key, code string) (string, string, error) {
	jsonstr := "{\"consumer_key\": \"" + key + "\", \"code\": \"" + code + "\"}"
	log.Printf("Info: %s", jsonstr)

	req, err := p.newReq("v3/oauth/authorize", jsonstr)
	if err != nil {
		return "", "", err
	}

	resp, err := p.c.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("Status code was not 200. got=%d", resp.StatusCode)
	}

	var r io.Reader = resp.Body
	//r = io.TeeReader(r, os.Stderr)

	var ra resAuth
	if err := json.NewDecoder(r).Decode(&ra); err != nil {
		return "", "", err
	}
	return ra.Username, ra.AccessToken, nil
}

func (p *pocket) GetList(key, token string) error {
	jsonstr := "{\"consumer_key\": \"" + key + "\", \"access_token\": \"" + token + "\", \"contentType\": \"article\"}"
	log.Printf("Info: %s", jsonstr)

	req, err := p.newReq("v3/get", jsonstr)

	if err != nil {
		return err
	}

	resp, err := p.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code was not 200. got=%d", resp.StatusCode)
	}

	var r io.Reader = resp.Body
	r = io.TeeReader(r, os.Stderr)

	var b interface{}
	if err := json.NewDecoder(r).Decode(&b); err != nil {
		return err
	}
	return nil
}
