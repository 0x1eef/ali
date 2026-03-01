package openai

type Completion struct {
	ID                string   `json:"id"`
	Choices           []Choice `json:"choices"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Object            string   `json:"object"`
	ServiceTier       string   `json:"service_tier,omitempty"`
	SystemFingerprint string   `json:"system_fingerprint,omitempty"`
	Usage             *Usage   `json:"usage,omitempty"`
}

type Choice struct {
	FinishReason string         `json:"finish_reason"`
	Index        int            `json:"index"`
	Logprobs     *ChatLogprobs  `json:"logprobs"`
	Message      *ChoiceMessage `json:"message,omitempty"`
}

type ChoiceMessage struct {
	Role        string `json:"role"`
	Content     string `json:"content"`
	Refusal     any    `json:"refusal,omitempty"`
	Annotations []any  `json:"annotations,omitempty"`
}

type ChatLogprobs struct {
	Content []TokenLogprob `json:"content,omitempty"`
	Refusal []TokenLogprob `json:"refusal,omitempty"`
}

type Usage struct {
	InputTokens   int           `json:"prompt_tokens"`
	OutputTokens  int           `json:"completion_tokens"`
	TotalTokens   int           `json:"total_tokens"`
	InputDetails  *TokenDetails `json:"prompt_tokens_details,omitempty"`
	OutputDetails *TokenDetails `json:"completion_tokens_details,omitempty"`
}

type TokenDetails struct {
	CachedTokens             int `json:"cached_tokens,omitempty"`
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens,omitempty"`
	AudioTokens              int `json:"audio_tokens,omitempty"`
	ReasoningTokens          int `json:"reasoning_tokens,omitempty"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens,omitempty"`
}

type TokenLogprob struct {
	Token       string       `json:"token"`
	Logprob     float64      `json:"logprob"`
	Bytes       []int        `json:"bytes,omitempty"`
	TopLogprobs []TopLogprob `json:"top_logprobs,omitempty"`
}

type TopLogprob struct {
	Token   string  `json:"token"`
	Logprob float64 `json:"logprob"`
	Bytes   []int   `json:"bytes,omitempty"`
}
