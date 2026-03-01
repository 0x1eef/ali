package ali

type ProviderName string
type Params map[string]any

const (
	OpenAI    ProviderName = "OpenAI"
	Anthropic ProviderName = "Anthropic"
	Gemini    ProviderName = "Gemini"
)

type Message struct {
	Role string `json:"-"`
	Text string `json:"-"`
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

type Provider interface {
	Name() ProviderName
	ApplyDefaults(*CompletionConfig) error
	Complete(options ...func(*CompletionConfig)) (Completion, error)
}
