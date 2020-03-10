package qiita

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func NewClient(c *http.Client) *Client {
	return &Client{
		c,
	}
}

type User struct {
	// omit unused fields.
	ID       string `json:"id"`
	Location string `json:"location"`
}

func (c *Client) FetchUser(id string) (*User, error) {
	req, err := http.NewRequest("GET", "https://qiita.com/api/v2/users/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	u := &User{}
	if err := json.Unmarshal(body, u); err != nil {
		return nil, err

	}
	return u, nil
}
