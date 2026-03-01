package session

import (
	"github.com/0x1eef/ali"
)

type Session struct {
	Provider ali.Provider `json:"-"`
	Options []func(*ali.CompletionConfig) `json:"-"`
	Messages []ali.Message
}

func New(provider ali.Provider, options ...func(*ali.CompletionConfig)) (*Session, error) {
	s := Session{Provider: provider, Options: options}
	return &s, nil
}

func (ses *Session) Talk(options ...func(*ali.CompletionConfig)) (ali.Completion, error) {
	allOptions := append(ses.Options, options...)
	allOptions = append(allOptions, ali.WithMessages(ses.Messages))
	completion, err := ses.Provider.Complete(allOptions...)
	if err != nil {
		return nil, err
	}
	ses.Messages = completion.Thread()
	return completion, err
}
