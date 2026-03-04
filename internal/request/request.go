package request

import (
	"io"
	"net/http"

	"github.com/0x1eef/ali"
)

func Post(options ...func(*config)) (*http.Response, error) {
	cfg := new(config)
	for _, set := range options {
		set(cfg)
	}
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	url, err := cfg.url()
	if err != nil {
		return nil, err
	}
	req, err := cfg.newRequest("POST", url)
	if err != nil {
		return nil, err
	}
	res, err := cfg.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		return nil, ali.ResponseError{
			StatusCode: res.StatusCode,
			Body:       body,
		}
	}
	return res, nil
}
