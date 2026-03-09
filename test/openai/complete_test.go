package openai_test

import (
	"os"
	"testing"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/openai"
	"github.com/0x1eef/ali/test/assert"
	"github.com/0x1eef/ali/test/vcr"
)

func TestComplete(t *testing.T) {
	token := os.Getenv("OPENAI_SECRET")
	if token == "" {
		t.Fatalf("OPENAI_SECRET is required")
	}

	client, stop := vcr.NewHTTPClient(t, "complete")
	defer stop()

	p, err := openai.New(openai.WithToken(token))
	assert.AssertNil(t, err)

	comp, err := p.Complete(
		ali.WithText("Reply with this exact string: 'pong'"),
		ali.WithClient(client),
	)
	assert.AssertNil(t, err)

	text, err := comp.Text()
	assert.AssertNil(t, err)
	assert.AssertEqual(t, text, "pong")
}
