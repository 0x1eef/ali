package gemini

import (
	"github.com/0x1eef/ali"
)

type Message struct {
	Role  string `json:"role,omitempty"`
	Parts []Part `json:"parts"`
}

type Part struct {
	Text             string            `json:"text,omitempty"`
	Thought          bool              `json:"thought,omitempty"`
	InlineData       *Blob             `json:"inlineData,omitempty"`
	FileData         *FileData         `json:"fileData,omitempty"`
	FunctionCall     *FunctionCall     `json:"functionCall,omitempty"`
	FunctionResponse *FunctionResponse `json:"functionResponse,omitempty"`
}

func toProviderMessages(cfg *ali.CompletionConfig) []Message {
	var (
		messages = make([]Message, 0, len(cfg.Messages)+1)
		message  = Message{Role: cfg.Role}
		parts    = []Part{}
	)
	for _, item := range cfg.Messages {
		parts := []Part{{Text: item.Text}}
		messages = append(messages, Message{Role: item.Role, Parts: parts})
	}
	if cfg.Prompt != "" {
		parts = append(parts, Part{Text: cfg.Prompt})
	}
	message.Parts = parts
	return append(messages, message)
}

func fromProviderMessages(completion *Completion) []ali.Message {
	msgs := make([]ali.Message, 0, len(completion.Candidates))
	msg := ali.Message{}
	for _, candidate := range completion.Candidates {
		msg.Role = candidate.Content.Role
		for _, part := range candidate.Content.Parts {
			if part.Text != "" {
				msg.Text = part.Text
			}
		}
		msgs = append(msgs, msg)
	}
	return msgs
}
