package openai

import (
	"github.com/0x1eef/ali"
)

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

func toProviderMessages(cfg *ali.CompletionConfig) []Message {
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
		contents = append(contents, Content{Type: "image_url", ImageUrl: url})
	}
	message.Content = contents
	return append(messages, message)
}

func fromProviderMessages(completion *Completion) []ali.Message {
	msgs := make([]ali.Message, 0, len(completion.Choices))
	for _, choice := range completion.Choices {
		msg := ali.Message{
			Role: choice.Message.Role,
			Text: choice.Message.Content,
		}
		msgs = append(msgs, msg)
	}
	return msgs
}
