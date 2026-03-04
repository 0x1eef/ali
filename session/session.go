package session

import "github.com/0x1eef/ali"

// Session keeps message history and default options for repeated provider calls.
type Session struct {
	Provider ali.Provider                  `json:"-"`
	Options  []func(*ali.CompletionConfig) `json:"-"`
	Messages []ali.Message
}

// New returns a Session that uses provider and applies options to every Talk call.
func New(provider ali.Provider, options ...func(*ali.CompletionConfig)) (*Session, error) {
	s := Session{Provider: provider, Options: options}
	return &s, nil
}

// Talk sends a completion request with session history and stores the returned thread.
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
