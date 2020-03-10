package qiita

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{
		&http.Client{},
	}
}

type User struct {
	Description       string `json:"description"`
	FacebookID        string `json:"facebook_id"`
	FolloweesCount    int    `json:"followees_count"`
	FollowersCount    int    `json:"followers_count"`
	GithubLoginName   string `json:"github_login_name"`
	ID                string `json:"id"`
	ItemsCount        int    `json:"items_count"`
	LinkedinID        string `json:"linkedin_id"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	Organization      string `json:"organization"`
	PermanentID       int    `json:"permanent_id"`
	ProfileImageURL   string `json:"profile_image_url"`
	TeamOnly          bool   `json:"team_only"`
	TwitterScreenName string `json:"twitter_screen_name"`
	WebsiteURL        string `json:"website_url"`
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
