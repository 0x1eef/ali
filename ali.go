package ali

import (
	"fmt"
	"io"
)

type ProviderName string
type Params map[string]any

const (
	OpenAI    ProviderName = "OpenAI"
	Anthropic ProviderName = "Anthropic"
	Gemini    ProviderName = "Gemini"
)

var (
	ErrNotImplemented = fmt.Errorf("feature is not implemented")
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
	// This can be one of the following:
	// 1) openai.Completion
	// 2) anthropic.Completion
	Raw() any
}

type Images interface {
	Create(options ...func(*ImageConfig)) ([]io.Reader, error)
}

type Provider interface {
	Name() ProviderName
	ApplyDefaults(*CompletionConfig) error
	Complete(options ...func(*CompletionConfig)) (Completion, error)
	Images() Images
}
