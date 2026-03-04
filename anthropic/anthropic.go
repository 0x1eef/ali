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
	token  string       `json:"-"`
	host   string       `json:"-"`
	client *http.Client `json:"-"`
}

func (provider *Anthropic) Name() ali.ProviderName {
	return ali.Anthropic
}

func New(options ...func(o *Anthropic)) (*Anthropic, error) {
	provider := Anthropic{host: "api.anthropic.com", client: &http.Client{}}
	for _, set := range options {
		set(&provider)
	}
	if provider.token == "" {
		return nil, fmt.Errorf("token is required")
	}
	return &provider, nil
}

func (provider *Anthropic) Complete(options ...func(cfg *ali.CompletionConfig)) (ali.Completion, error) {
	var comp Completion
	var err error
	cfg := ali.CompletionConfig{Provider: provider}
	if err := cfg.ApplyDefaults(options...); err != nil {
		return nil, err
	}
	payload := struct {
		Model     string    `json:"model"`
		Messages  []Message `json:"messages"`
		MaxTokens int       `json:"max_tokens"`
	}{
		Model:     cfg.Model,
		Messages:  toProviderMessages(&cfg),
		MaxTokens: cfg.MaxTokens,
	}
	body, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}
	res, err := request.Post(
		request.WithHost(provider.host),
		request.WithPath("/v1/messages"),
		request.WithBody(bytes.NewReader(body)),
		request.WithParams(cfg.Params),
		request.WithClient(provider.client),
		request.WithSetup(func(req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("x-api-key", provider.token)
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

func (provider *Anthropic) ApplyDefaults(cfg *ali.CompletionConfig) error {
	if cfg.Model == "" {
		cfg.Model = "claude-sonnet-4-20250514"
	}
	if cfg.MaxTokens == 0 {
		cfg.MaxTokens = 1024
	}
	return nil
}

func (provider *Anthropic) Images() ali.Images {
	return Images{}
}
