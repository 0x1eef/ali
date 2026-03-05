package anthropic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/internal/request"
)

type Anthropic struct {
	name   ali.ProviderName `json:"-"`
	token  string           `json:"-"`
	host   string           `json:"-"`
	client *http.Client     `json:"-"`
}

func (ant *Anthropic) Name() ali.ProviderName {
	return ant.name
}

func New(options ...func(o *Anthropic)) (*Anthropic, error) {
	provider := Anthropic{name: ali.Anthropic, host: "api.anthropic.com", client: &http.Client{}}
	for _, set := range options {
		set(&provider)
	}
	if provider.token == "" {
		return nil, fmt.Errorf("token is required")
	}
	return &provider, nil
}

func (ant *Anthropic) Complete(options ...func(cfg *ali.CompletionConfig)) (ali.Completion, error) {
	var comp Completion
	var err error
	cfg := ali.CompletionConfig{Provider: ant}
	if err := cfg.ApplyDefaults(options...); err != nil {
		return nil, err
	}
	params, err := ant.build(&cfg)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	res, err := request.Post(
		request.WithHost(ant.host),
		request.WithPath("/v1/messages"),
		request.WithBody(bytes.NewReader(body)),
		request.WithClient(ant.client),
		request.WithSetup(func(req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("x-api-key", ant.token)
			req.Header.Add("anthropic-version", "2023-06-01")
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

func (ant *Anthropic) ApplyDefaults(cfg *ali.CompletionConfig) error {
	if cfg.Model == "" {
		cfg.Model = "claude-sonnet-4-20250514"
	}
	if cfg.MaxTokens == 0 {
		cfg.MaxTokens = 1024
	}
	return nil
}

func (ant *Anthropic) build(cfg *ali.CompletionConfig) (ali.Params, error) {
	mesgs, err := toProviderMessages(cfg)
	if err != nil {
		return nil, err
	}
	params := ali.Params{
		"model":      cfg.Model,
		"messages":   mesgs,
		"max_tokens": cfg.MaxTokens,
	}
	for k, v := range cfg.Params {
		params[k] = v
	}
	return params, nil
}
