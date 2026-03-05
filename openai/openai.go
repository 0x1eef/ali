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
	name  ali.ProviderName
	Token string
	Host  string
}

func (oai *OpenAI) Name() ali.ProviderName {
	return oai.name
}

func New(options ...func(o *OpenAI)) (*OpenAI, error) {
	oai := OpenAI{name: ali.OpenAI, Host: "api.openai.com"}
	for _, set := range options {
		set(&oai)
	}
	if oai.Token == "" {
		return nil, fmt.Errorf("token is required")
	}
	return &oai, nil
}

func (oai *OpenAI) Complete(options ...func(*ali.CompletionConfig)) (ali.Completion, error) {
	var comp Completion
	var err error
	cfg := ali.CompletionConfig{Provider: oai}
	if err = cfg.ApplyDefaults(options...); err != nil {
		return nil, err
	}
	params, err := oai.build(&cfg)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	res, err := request.Post(
		request.WithHost(oai.Host),
		request.WithPath("/v1/chat/completions"),
		request.WithBody(bytes.NewReader(body)),
		request.WithClient(cfg.Client),
		request.WithSetup(func(req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", oai.Token))
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

func (oai *OpenAI) ApplyDefaults(cfg *ali.CompletionConfig) error {
	if cfg.Model == "" {
		cfg.Model = "gpt-4.1"
	}
	return nil
}

func (oai *OpenAI) Images() ali.Images {
	return Images{provider: oai}
}

func (oai *OpenAI) build(cfg *ali.CompletionConfig) (ali.Params, error) {
	mesgs, err := toProviderMessages(cfg)
	if err != nil {
		return nil, err
	}
	params := ali.Params{
		"model":    cfg.Model,
		"messages": mesgs,
	}
	if cfg.MaxTokens != 0 {
		params["max_completion_tokens"] = cfg.MaxTokens
	}
	for k, v := range cfg.Params {
		params[k] = v
	}
	return params, nil
}
