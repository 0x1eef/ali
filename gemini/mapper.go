package gemini

import (
	"encoding/base64"
	"io"
	"os"

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

func toProviderMessages(cfg *ali.CompletionConfig) ([]Message, error) {
	var (
		messages = make([]Message, 0, len(cfg.Messages)+1)
		message  = Message{Role: cfg.Role}
		parts    = []Part{}
	)
	for _, item := range cfg.Messages {
		parts := []Part{{Text: item.Text}}
		messages = append(messages, Message{Role: item.Role, Parts: parts})
	}
	for _, text := range cfg.Texts {
		parts = append(parts, Part{Text: text})
	}
	for _, url := range cfg.ImageUrls {
		fileData := FileData{FileURI: url}
		parts = append(parts, Part{FileData: &fileData})
	}
	for _, pdf := range cfg.Pdfs {
		part, err := fileToContent(pdf, "application/pdf")
		if err != nil {
			return nil, err
		}
		parts = append(parts, part)
	}
	message.Parts = parts
	return append(messages, message), nil
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

func fileToContent(file, kind string) (Part, error) {
	f, err := os.Open(file)
	if err != nil {
		return Part{}, err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return Part{}, err
	}
	b64 := base64.StdEncoding.EncodeToString(b)
	inlineData := Blob{MIMEType: kind, Data: b64}
	return Part{InlineData: &inlineData}, nil
}
