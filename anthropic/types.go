package anthropic

type Completion struct {
	ID           string         `json:"id"`
	Container    *Container     `json:"container,omitempty"`
	Content      []ContentBlock `json:"content"`
	Model        string         `json:"model"`
	Role         string         `json:"role"`
	StopReason   string         `json:"stop_reason"`
	StopSequence string         `json:"stop_sequence,omitempty"`
	Type         string         `json:"type"`
	Usage        Usage          `json:"usage"`
}

type Container struct {
	ID        string `json:"id"`
	ExpiresAt string `json:"expires_at"`
}

// ContentBlock intentionally models common fields across multiple block variants.
// For unknown or rapidly changing shapes, Content carries raw JSON-compatible data.
type ContentBlock struct {
	Type string `json:"type"`

	// Text and citations
	Text      string     `json:"text,omitempty"`
	Citations []Citation `json:"citations,omitempty"`

	// Thinking / redacted thinking
	Thinking  string `json:"thinking,omitempty"`
	Signature string `json:"signature,omitempty"`
	Data      string `json:"data,omitempty"`

	// Tool use / server tool use
	ID     string                 `json:"id,omitempty"`
	Name   string                 `json:"name,omitempty"`
	Input  map[string]interface{} `json:"input,omitempty"`
	Caller *Caller                `json:"caller,omitempty"`

	// Tool results
	ToolUseID string      `json:"tool_use_id,omitempty"`
	Content   interface{} `json:"content,omitempty"`

	// Container upload and other file-like blocks
	FileID string `json:"file_id,omitempty"`
}

type Caller struct {
	Type   string `json:"type"`
	ToolID string `json:"tool_id,omitempty"`
}

type Citation struct {
	Type string `json:"type"`

	CitedText     string `json:"cited_text,omitempty"`
	DocumentIndex int    `json:"document_index,omitempty"`
	DocumentTitle string `json:"document_title,omitempty"`
	FileID        string `json:"file_id,omitempty"`

	// char_location
	StartCharIndex int `json:"start_char_index,omitempty"`
	EndCharIndex   int `json:"end_char_index,omitempty"`

	// page_location
	StartPageNumber int `json:"start_page_number,omitempty"`
	EndPageNumber   int `json:"end_page_number,omitempty"`

	// content_block_location
	StartBlockIndex int `json:"start_block_index,omitempty"`
	EndBlockIndex   int `json:"end_block_index,omitempty"`

	// web_search_result_location
	EncryptedIndex string `json:"encrypted_index,omitempty"`
	Title          string `json:"title,omitempty"`
	URL            string `json:"url,omitempty"`

	// search_result_location
	SearchResultIndex int    `json:"search_result_index,omitempty"`
	Source            string `json:"source,omitempty"`
}

type Usage struct {
	CacheCreationInputTokens int            `json:"cache_creation_input_tokens,omitempty"`
	CacheReadInputTokens     int            `json:"cache_read_input_tokens,omitempty"`
	InputTokens              int            `json:"input_tokens"`
	OutputTokens             int            `json:"output_tokens"`
	InferenceGeo             string         `json:"inference_geo,omitempty"`
	ServiceTier              string         `json:"service_tier,omitempty"`
	CacheCreation            *CacheCreation `json:"cache_creation,omitempty"`
	ServerToolUse            *ServerToolUse `json:"server_tool_use,omitempty"`
}

type CacheCreation struct {
	Ephemeral1hInputTokens int `json:"ephemeral_1h_input_tokens,omitempty"`
	Ephemeral5mInputTokens int `json:"ephemeral_5m_input_tokens,omitempty"`
}

type ServerToolUse struct {
	WebFetchRequests  int `json:"web_fetch_requests,omitempty"`
	WebSearchRequests int `json:"web_search_requests,omitempty"`
}
