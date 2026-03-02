package provider

import (
	"fmt"
	"os"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/anthropic"
	"github.com/0x1eef/ali/gemini"
	"github.com/0x1eef/ali/openai"
)

const (
	OpenAI    = ali.OpenAI
	Anthropic = ali.Anthropic
	Gemini    = ali.Gemini
)

func New(providerName ali.ProviderName) (ali.Provider, error) {
	switch providerName {
	case ali.Gemini:
		return gemini.New(
			gemini.WithToken(os.Getenv("GEMINI_SECRET")),
		)
	case ali.OpenAI:
		return openai.New(
			openai.WithToken(os.Getenv("OPENAI_SECRET")),
		)
	case ali.Anthropic:
		return anthropic.New(
			anthropic.WithToken(os.Getenv("ANTHROPIC_SECRET")),
		)
	default:
		return nil, fmt.Errorf("unknown provider: %s", providerName)
	}
}
