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
	name  ali.ProviderName `json:"-"`
	Token string           `json:"-"`
	Host  string           `json:"-"`
}

func (gem *Gemini) Name() ali.ProviderName {
	return gem.name
}

func New(options ...func(o *Gemini)) (*Gemini, error) {
	provider := Gemini{name: ali.Gemini, Host: "generativelanguage.googleapis.com"}
	for _, set := range options {
		set(&provider)
	}
	if provider.Token == "" {
		return nil, fmt.Errorf("token is required")
	}
	return &provider, nil
}

func (gem *Gemini) Complete(options ...func(cfg *ali.CompletionConfig)) (ali.Completion, error) {
	var comp Completion
	var err error
	cfg := ali.CompletionConfig{Provider: gem}
	if err := cfg.ApplyDefaults(options...); err != nil {
		return nil, err
	}
	params, err := gem.build(&cfg)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	res, err := request.Post(
		request.WithHost(gem.Host),
		request.WithPath(fmt.Sprintf("/v1beta/models/%s:generateContent?key=%s", cfg.Model, gem.Token)),
		request.WithBody(bytes.NewReader(body)),
		request.WithClient(cfg.Client),
		request.WithContext(cfg.Ctx),
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

func (gem *Gemini) ApplyDefaults(cfg *ali.CompletionConfig) error {
	if cfg.Role == "" {
		cfg.Role = "user"
	}
	if cfg.Model == "" {
		cfg.Model = "gemini-2.5-flash"
	}
	return nil
}

func (gem *Gemini) Images() ali.Images {
	return Images{provider: gem}
}

func (gem *Gemini) build(cfg *ali.CompletionConfig) (ali.Params, error) {
	mesgs, err := toProviderMessages(cfg)
	if err != nil {
		return nil, err
	}
	params := ali.Params{
		"contents": mesgs,
	}
	for k, v := range cfg.Params {
		params[k] = v
	}
	return params, nil
}
