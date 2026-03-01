package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/internal/request"
)

type Gemini struct {
	name   ali.ProviderName `json:"-"`
	token  string           `json:"-"`
	host   string           `json:"-"`
	client *http.Client     `json:"-"`
}

func New(options ...func(o *Gemini)) (*Gemini, error) {
	provider := Gemini{host: "generativelanguage.googleapis.com", client: &http.Client{}}
	for _, set := range options {
		set(&provider)
	}
	if provider.token == "" {
		return nil, fmt.Errorf("token is required")
	}
	return &provider, nil
}

func (provider *Gemini) Complete(options ...func(cfg *ali.CompletionConfig)) (ali.Completion, error) {
	var comp Completion
	var err error
	cfg := ali.CompletionConfig{Provider: provider}
	for _, set := range options {
		set(&cfg)
	}
	if err := cfg.ApplyDefaults(); err != nil {
		return nil, err
	}
	payload := struct {
		Contents []Message `json:"contents"`
	}{
		Contents: toProviderMessages(&cfg),
	}
	body, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}
	if cfg.Params == nil {
		cfg.Params = make(ali.Params)
	}
	cfg.Params["key"] = provider.token
	res, err := request.Post(
		request.WithHost(provider.host),
		request.WithPath(fmt.Sprintf("/v1/beta/models/%s/generateContent", cfg.Model)),
		request.WithBody(bytes.NewReader(body)),
		request.WithParams(cfg.Params),
		request.WithClient(provider.client),
		request.WithSetup(func(req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&comp); err != nil {
		return nil, err
	}
	return CompletionAdapter{completion: &comp, thread: cfg.Messages}, nil
}

func (provider *Gemini) Name() ali.ProviderName {
	return provider.name
}

func (provider *Gemini) ApplyDefaults(cfg *ali.CompletionConfig) error {
	if cfg.Role != "" {
		cfg.Role = "user"
	}
	if cfg.Model != "" {
		cfg.Model = "gemini-2.5-flash"
	}
	return nil
}
