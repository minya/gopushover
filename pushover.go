package gopushover

import (
	//"errors"
	"encoding/json"
	"fmt"
	"github.com/minya/goutils/web"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PushoverResult struct represents response from pushover
type PushoverResult struct {
	Status  int      `json:"status"`
	Request string   `json:"request"`
	Errors  []string `json:"errors"`
}

// SendMessage sends message to pushover
func SendMessage(token string, user string, msg string) (res *PushoverResult, e error) {
	client := http.Client{
		Transport: web.DefaultTransport(1000),
	}

	form := url.Values{
		"message": {msg},
		"token":   {token},
		"user":    {user},
	}

	resp, err := client.PostForm("https://api.pushover.net/1/messages.json", form)
	if err != nil {
		return nil, err
	}

	fmt.Printf("resp from pushover: %v\n", resp.StatusCode)

	bodyBin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := new(PushoverResult)
	err = json.Unmarshal(bodyBin, result)
	if err != nil {
		return nil, NewPushoverError(string(bodyBin))
	}

	return result, nil
}

type PushoverError struct {
	msg string
}

func (e *PushoverError) Error() string {
	return e.msg
}

func NewPushoverError(msg string) error {
	e := new(PushoverError)
	e.msg = msg
	return e
}
