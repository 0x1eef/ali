package session

import (
	"github.com/0x1eef/ali"
)

type Session struct {
	Provider ali.Provider `json:"-"`
	Messages []ali.Message
}

func New(provider ali.Provider) (*Session, error) {
	s := Session{Provider: provider}
	return &s, nil
}

func (ses *Session) Talk(options ...func(*ali.CompletionConfig)) (ali.Completion, error) {
	options = append(options, ali.WithMessages(ses.Messages))
	completion, err := ses.Provider.Complete(options...)
	if err != nil {
		return nil, err
	}
	ses.Messages = completion.Thread()
	return completion, err
}
