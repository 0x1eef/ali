package openai

import (
	"fmt"

	"github.com/0x1eef/ali"
)

type CompletionAdapter struct {
	completion *Completion
	thread     []ali.Message
}

func (ca CompletionAdapter) Raw() any {
	return ca.completion
}

func (ca CompletionAdapter) Messages() []ali.Message {
	return fromProviderMessages(ca.completion)
}

func (ca CompletionAdapter) Usage() ali.Usage {
	return ali.Usage{
		InputTokens:  ca.completion.Usage.InputTokens,
		OutputTokens: ca.completion.Usage.OutputTokens,
		TotalTokens:  ca.completion.Usage.TotalTokens,
	}
}

func (ca CompletionAdapter) Text() (string, error) {
	if choices := ca.Messages(); len(choices) > 0 {
		return choices[0].Text, nil
	} else {
		return "", fmt.Errorf("no suitable choices found")
	}
}

func (ca CompletionAdapter) Thread() []ali.Message {
	thr := make([]ali.Message, 0, len(ca.thread)+1)
	thr = append(thr, ca.thread...)
	if msgs := ca.Messages(); len(msgs) > 0 {
		thr = append(thr, msgs[0])
	}
	return thr
}
