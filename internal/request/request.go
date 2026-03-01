package request

import (
	"net/http"

	"github.com/0x1eef/ali/errors"
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
		return nil, errors.ResponseError{Response: res}
	}
	return res, nil
}
