package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/internal/request"
)

type OpenAI struct {
	name   ali.ProviderName
	token  string
	host   string
	client *http.Client
}

func (provider *OpenAI) Name() ali.ProviderName {
	return provider.name
}

func New(options ...func(o *OpenAI)) (*OpenAI, error) {
	provider := OpenAI{name: ali.OpenAI, host: "api.openai.com", client: &http.Client{}}
	for _, set := range options {
		set(&provider)
	}
	if provider.token == "" {
		return nil, fmt.Errorf("token is required")
	}
	return &provider, nil
}

func (provider *OpenAI) Complete(options ...func(*ali.CompletionConfig)) (ali.Completion, error) {
	var comp Completion
	var err error
	cfg := ali.CompletionConfig{Provider: provider}
	if err = cfg.ApplyDefaults(options...); err != nil {
		return nil, err
	}
	payload := struct {
		Model     string    `json:"model"`
		Messages  []Message `json:"messages"`
		MaxTokens int       `json:"max_completion_tokens,omitempty"`
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
		request.WithPath("/v1/chat/completions"),
		request.WithBody(bytes.NewReader(body)),
		request.WithParams(cfg.Params),
		request.WithClient(provider.client),
		request.WithSetup(func(req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", provider.token))
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

func (provider *OpenAI) ApplyDefaults(cfg *ali.CompletionConfig) error {
	if cfg.Model == "" {
		cfg.Model = "gpt-4.1"
	}
	return nil
}

func (provider *OpenAI) Images() (ali.Images) {
	return Images{}
}
