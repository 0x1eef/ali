package ali

import (
	"context"
)

type CompletionConfig struct {
	Provider  Provider        `json:"-"`
	Prompt    string          `json:"-"`
	Role      string          `json:"-"`
	Params    Params          `json:"-"`
	Ctx       context.Context `json:"-"`
	Model     string          `json:"-"`
	Messages  []Message       `json:"-"`
	MaxTokens int             `json:"-"`
}

// WithPrompt sets the prompt text for a request.
func WithPrompt(prompt string) func(r *CompletionConfig) {
	return func(r *CompletionConfig) {
		r.Prompt = prompt
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
	return cfg.Provider.ApplyDefaults(cfg)
}
