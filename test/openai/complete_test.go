package openai_test

import (
	"testing"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"

	"github.com/0x1eef/ali/test/assert"
	"github.com/0x1eef/ali/test/vcr"
)

func TestComplete(t *testing.T) {
	p, err := provider.New(ali.OpenAI)
	assert.AssertNil(t, err)

	client, stop := vcr.NewHTTPClient(t, "complete")
	defer stop()

	comp, err := p.Complete(
		ali.WithText("Reply with this exact string: 'pong'"),
		ali.WithClient(client),
	)
	assert.AssertNil(t, err)

	text, err := comp.Text()
	assert.AssertNil(t, err)
	assert.AssertEqual(t, text, "pong")
}
