package ali

import (
	"io"
)

type ProviderName string
type Params map[string]any

const (
	OpenAI    ProviderName = "OpenAI"
	Anthropic ProviderName = "Anthropic"
	Gemini    ProviderName = "Gemini"
)

type Message struct {
	Role string `json:"role"`
	Text string `json:"content"`
}

type Completion interface {
	InputTokens() int
	OutputTokens() int
	TotalTokens() int
	Text() (string, error)
	Messages() []Message
	Thread() []Message
	// Raw returns the provider-specific completion payload (eg openai.Completion).
	Raw() any
}

type Images interface {
	Create(options ...func(*ImageConfig)) ([]io.Reader, error)
}

type Provider interface {
	Name() ProviderName
	ApplyDefaults(*CompletionConfig) error
	Complete(options ...func(*CompletionConfig)) (Completion, error)
}

type ImageProvider interface {
	Images() Images
}
