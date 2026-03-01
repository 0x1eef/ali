package main

import (
	"fmt"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
)

func main() {
	provider, err := provider.Select(ali.OpenAI)
	if err != nil {
		panic(err)
	}

	completion, err := provider.Complete(
		ali.WithPrompt("Hello from #golang :)"),
	)
	if err != nil {
		panic(err)
	}

	text, err := completion.Text()
	if err != nil {
		panic(err)
	}
	fmt.Printf("LLM says: %s\n", text)
}
