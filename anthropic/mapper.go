package anthropic

import (
	"encoding/base64"
	"io"
	"os"

	"github.com/0x1eef/ali"
)

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type   string `json:"type"`
	Text   string `json:"text,omitempty"`
	Source Source `json:"source,omitempty"`
}

type Source struct {
	Type      string `json:"type"`
	Url       string `json:"url,omitempty"`
	MediaType string `json:"media_type,omitempty"`
	Data      string `json:"data,omitempty"`
}

func toProviderMessages(cfg *ali.CompletionConfig) ([]Message, error) {
	var (
		messages = make([]Message, 0, len(cfg.Messages)+1)
		message  = Message{Role: cfg.Role}
		contents = []Content{}
	)
	for _, item := range cfg.Messages {
		content := []Content{{Type: "text", Text: item.Text}}
		messages = append(messages, Message{Role: item.Role, Content: content})
	}
	for _, text := range cfg.Texts {
		contents = append(contents, Content{Type: "text", Text: text})
	}
	for _, url := range cfg.ImageUrls {
		source := Source{Type: "url", Url: url}
		content := Content{Type: "image", Source: source}
		contents = append(contents, content)
	}
	for _, pdf := range cfg.Pdfs {
		content, err := fileToContent(pdf, "application/pdf")
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	message.Content = contents
	return append(messages, message), nil
}

func fromProviderMessages(completion *Completion) []ali.Message {
	msgs := make([]ali.Message, 0, len(completion.Content))
	for _, contentBlock := range completion.Content {
		msg := ali.Message{
			Role: completion.Role,
			Text: contentBlock.Text,
		}
		msgs = append(msgs, msg)
	}
	return msgs
}

func fileToContent(file, kind string) (Content, error) {
	f, err := os.Open(file)
	if err != nil {
		return Content{}, err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return Content{}, err
	}
	b64 := base64.StdEncoding.EncodeToString(b)
	return Content{
		Type:   "document",
		Source: Source{Type: "base64", MediaType: kind, Data: b64},
	}, nil
}
