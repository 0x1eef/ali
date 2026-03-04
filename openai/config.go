package openai

import (
	"context"

	"github.com/0x1eef/ali"
)

// WithToken sets the OpenAI API token.
func WithToken(token string) func(o *OpenAI) {
	return func(o *OpenAI) {
		o.Token = token
	}
}

// WithHost sets the OpenAI API host.
func WithHost(host string) func(o *OpenAI) {
	return func(o *OpenAI) {
		o.Host = host
	}
}

// WithPrompt sets the user prompt text.
func WithPrompt(prompt string) func(r *ali.CompletionConfig) {
	return ali.WithPrompt(prompt)
}

// WithRole sets the role for the prompt message.
func WithRole(role string) func(r *ali.CompletionConfig) {
	return ali.WithRole(role)
}

// WithModel sets the model identifier.
func WithModel(model string) func(r *ali.CompletionConfig) {
	return ali.WithModel(model)
}

// WithParams sets request query parameters.
func WithParams(params ali.Params) func(r *ali.CompletionConfig) {
	return ali.WithParams(params)
}

// WithContext sets the request context.
func WithContext(ctx context.Context) func(r *ali.CompletionConfig) {
	return ali.WithContext(ctx)
}

// WithMessages appends prior conversation messages.
func WithMessages(msgs []ali.Message) func(r *ali.CompletionConfig) {
	return ali.WithMessages(msgs)
}
