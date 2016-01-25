package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/naoina/toml"
)

// Slack has webhook parameters
type Slack struct {
	Payload
	URL string
}

// Payload includes `payload` slack parameters
type Payload struct {
	Channel   string `json:"channel"`
	UserName  string `json:"username" toml:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	LinkNames int    `json:"link_names"`
}

// NewSlack returns a initialized Slack struct by config.toml
func NewSlack() (*Slack, error) {
	var s = &Slack{
		Payload: Payload{},
		URL:     "",
	}

	buf, err := s.readConfigToml()
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(buf, s); err != nil {
		return nil, err
	}

	return s, nil
}

// ConfigPath is a file path of config.toml
const ConfigPath = "config/config.toml"

func (s *Slack) readConfigToml() (buf []byte, err error) {
	f, err := os.Open(ConfigPath)
	if err != nil {
		return
	}
	defer f.Close()

	buf, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}

	return
}

// Post submits post request to slack
func (s *Slack) Post() error {
	b, err := json.Marshal(s.Payload)
	if err != nil {
		return err
	}

	var v = url.Values{}
	v.Set("payload", string(b))

	res, err := http.PostForm(s.URL, v)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
