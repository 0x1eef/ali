package openai

import (
	"encoding/base64"
	"fmt"
	"os"
	"io"
	"path/filepath"

	"github.com/0x1eef/ali"
)

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type File struct {
	Filename string `json:"filename"`
	FileData string `json:"file_data"`
}

type Content struct {
	Type     string `json:"type"`
	Text     string `json:"text,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	File     File   `json:"file,omitempty"`
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
		contents = append(contents, Content{Type: "image_url", ImageUrl: url})
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
	base := filepath.Base(file)
	data := fmt.Sprintf("data:%s;base64,%s", kind, b64)
	return Content{
		Type: "file",
		File: File{Filename: base, FileData: data},
	}, nil
}
