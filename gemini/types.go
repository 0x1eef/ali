package gemini

type Completion struct {
	ResponseID     string          `json:"responseId"`
	Candidates     []Candidate     `json:"candidates,omitempty"`
	PromptFeedback *PromptFeedback `json:"promptFeedback,omitempty"`
	UsageMetadata  *UsageMetadata  `json:"usageMetadata,omitempty"`
	ModelVersion   string          `json:"modelVersion,omitempty"`
	CreateTime     string          `json:"createTime,omitempty"`
	ModelStatus    *ModelStatus    `json:"modelStatus,omitempty"`
}

type Candidate struct {
	Index            int               `json:"index,omitempty"`
	Content          Content           `json:"content"`
	FinishReason     string            `json:"finishReason,omitempty"`
	FinishMessage    string            `json:"finishMessage,omitempty"`
	SafetyRatings    []SafetyRating    `json:"safetyRatings,omitempty"`
	CitationMetadata *CitationMetadata `json:"citationMetadata,omitempty"`
	TokenCount       int               `json:"tokenCount,omitempty"`
	AvgLogprobs      float64           `json:"avgLogprobs,omitempty"`
	LogprobsResult   *LogprobsResult   `json:"logprobsResult,omitempty"`
}

type PromptFeedback struct {
	BlockReason        string         `json:"blockReason,omitempty"`
	BlockReasonMessage string         `json:"blockReasonMessage,omitempty"`
	SafetyRatings      []SafetyRating `json:"safetyRatings,omitempty"`
}

type UsageMetadata struct {
	PromptTokenCount          int                  `json:"promptTokenCount,omitempty"`
	CandidatesTokenCount      int                  `json:"candidatesTokenCount,omitempty"`
	TotalTokenCount           int                  `json:"totalTokenCount,omitempty"`
	CachedContentTokenCount   int                  `json:"cachedContentTokenCount,omitempty"`
	ToolUsePromptTokenCount   int                  `json:"toolUsePromptTokenCount,omitempty"`
	ThoughtsTokenCount        int                  `json:"thoughtsTokenCount,omitempty"`
	PromptTokensDetails       []ModalityTokenCount `json:"promptTokensDetails,omitempty"`
	CandidatesTokensDetails   []ModalityTokenCount `json:"candidatesTokensDetails,omitempty"`
	CacheTokensDetails        []ModalityTokenCount `json:"cacheTokensDetails,omitempty"`
	ToolUsePromptTokensDetail []ModalityTokenCount `json:"toolUsePromptTokensDetails,omitempty"`
}

type ModelStatus map[string]any

type Content struct {
	Role  string `json:"role,omitempty"`
	Parts []Part `json:"parts"`
}


type Blob struct {
	MIMEType string `json:"mimeType"`
	Data     string `json:"data"`
}

type FileData struct {
	MIMEType string `json:"mimeType,omitempty"`
	FileURI  string `json:"fileUri,omitempty"`
}

type FunctionCall struct {
	Name string         `json:"name"`
	Args map[string]any `json:"args,omitempty"`
}

type FunctionResponse struct {
	Name     string         `json:"name"`
	Response map[string]any `json:"response,omitempty"`
}

type SafetyRating struct {
	Category    string `json:"category,omitempty"`
	Probability string `json:"probability,omitempty"`
	Severity    string `json:"severity,omitempty"`
	Blocked     bool   `json:"blocked,omitempty"`
}

type CitationMetadata struct {
	CitationSources []CitationSource `json:"citationSources,omitempty"`
}

type CitationSource struct {
	StartIndex int    `json:"startIndex,omitempty"`
	EndIndex   int    `json:"endIndex,omitempty"`
	URI        string `json:"uri,omitempty"`
	License    string `json:"license,omitempty"`
}

type LogprobsResult struct {
	TopCandidates    []LogprobsCandidates `json:"topCandidates,omitempty"`
	ChosenCandidates []LogprobsCandidate  `json:"chosenCandidates,omitempty"`
}

type LogprobsCandidates struct {
	Candidates []LogprobsCandidate `json:"candidates,omitempty"`
}

type LogprobsCandidate struct {
	TokenID int     `json:"tokenId,omitempty"`
	Token   string  `json:"token,omitempty"`
	Logprob float64 `json:"logProbability,omitempty"`
}

type ModalityTokenCount struct {
	Modality   string `json:"modality,omitempty"`
	TokenCount int    `json:"tokenCount,omitempty"`
}
