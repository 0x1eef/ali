package request

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	neturl "net/url"
	"strconv"

	"github.com/0x1eef/ali"
)

type config struct {
	host   string
	path   string
	params ali.Params
	setup  func(req *http.Request) error
	body   *bytes.Reader
	client *http.Client
	ctx    context.Context
}

func WithParams(params ali.Params) func(cfg *config) {
	return func(cfg *config) {
		cfg.params = params
	}
}

func WithHost(host string) func(cfg *config) {
	return func(cfg *config) {
		cfg.host = host
	}
}

func WithPath(path string) func(cfg *config) {
	return func(cfg *config) {
		cfg.path = path
	}
}

func WithClient(client *http.Client) func(cfg *config) {
	return func(cfg *config) {
		cfg.client = client
	}
}

func WithBody(body *bytes.Reader) func(cfg *config) {
	return func(cfg *config) {
		cfg.body = body
	}
}

func WithSetup(setup func(req *http.Request) error) func(cfg *config) {
	return func(cfg *config) {
		cfg.setup = setup
	}
}

func WithContext(ctx context.Context) func(cfg *config) {
	return func(cfg *config) {
		cfg.ctx = ctx
	}
}

func (cfg config) validate() error {
	if cfg.host == "" {
		return fmt.Errorf("a host is required")
	}
	if cfg.path == "" {
		return fmt.Errorf("a path is required")
	}
	if cfg.setup == nil {
		return fmt.Errorf("a setup function is required")
	}
	if cfg.client == nil {
		return fmt.Errorf("a client is required")
	}
	return nil
}

func (cfg *config) url() (string, error) {
	rawUrl := fmt.Sprintf("https://%s%s", cfg.host, cfg.path)
	url, err := neturl.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	q := url.Query()
	for k, v := range cfg.params {
		if s, ok := v.(string); ok {
			q.Add(k, s)
		} else if d, ok := v.(int); ok {
			q.Add(k, strconv.Itoa(d))
		} else if d, ok := v.(int64); ok {
			q.Add(k, strconv.FormatInt(d, 10))
		} else if f, ok := v.(float32); ok {
			q.Add(k, strconv.FormatFloat(float64(f), 'f', -1, 32))
		} else if f, ok := v.(float64); ok {
			q.Add(k, strconv.FormatFloat(f, 'f', -1, 64))
		} else if b, ok := v.(bool); ok {
			if b {
				q.Add(k, "true")
			} else {
				q.Add(k, "false")
			}
		} else {
			return "", fmt.Errorf("don't know how to handle param %s", k)
		}
	}
	url.RawQuery = q.Encode()
	return url.String(), nil
}

func (cfg *config) newRequest(verb, url string) (*http.Request, error) {
	req, err := http.NewRequest(verb, url, cfg.body)
	if err != nil {
		return nil, err
	}
	if cfg.ctx != nil {
		req = req.WithContext(cfg.ctx)
	}
	if err := cfg.setup(req); err != nil {
		return nil, err
	}
	return req, nil
}
