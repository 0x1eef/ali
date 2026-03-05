package ali

import (
	"context"
	"net/http"
)

type CompletionConfig struct {
	Provider  Provider        `json:"-"`
	Texts     []string        `json:"-"`
	ImageUrls []string        `json:"-"`
	Pdfs      []string        `json:"-"`
	Role      string          `json:"-"`
	Params    Params          `json:"-"`
	Ctx       context.Context `json:"-"`
	Client    *http.Client    `json:"-"`
	Model     string          `json:"-"`
	Messages  []Message       `json:"-"`
	MaxTokens int             `json:"-"`
}

type ImageConfig struct {
	Prompt   string          `json:"-"`
	Quantity int             `json:"-"`
	Model    string          `json:"-"`
	Params   Params          `json:"-"`
	Ctx      context.Context `json:"-"`
	Client   *http.Client    `json:"-"`
}

// WithText sets the prompt text for a request.
func WithText(prompt string) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Texts = append(r.Texts, prompt)
	}
}

// WithImageUrl includes an image URL with a request.
func WithImageUrl(url string) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.ImageUrls = append(r.ImageUrls, url)
	}
}

// WithPdf includes a PDF with a request.
func WithPdf(pdf string) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Pdfs = append(r.Pdfs, pdf)
	}
}

// WithRole sets the role used for prompt-derived messages.
func WithRole(role string) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Role = role
	}
}

// WithModel sets the provider model identifier.
func WithModel(model string) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Model = model
	}
}

// WithParams sets provider query parameters.
func WithParams(params Params) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Params = params
	}
}

// WithContext sets the request context.
func WithContext(ctx context.Context) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Ctx = ctx
	}
}

// WithClient sets the http client.
func WithClient(client *http.Client) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Client = client
	}
}

// WithMessages appends existing conversation messages.
func WithMessages(msgs []Message) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Messages = append(r.Messages, msgs...)
	}
}

func (cfg *CompletionConfig) ApplyDefaults(options ...func(*CompletionConfig)) error {
	for _, set := range options {
		set(cfg)
	}
	if cfg.Role == "" {
		cfg.Role = "user"
	}
	if cfg.Client == nil {
		cfg.Client = &http.Client{}
	}
	return cfg.Provider.ApplyDefaults(cfg)
}
