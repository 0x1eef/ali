package main

import (
	"fmt"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
)

func main() {
	p, err := provider.New(ali.OpenAI)
	if err != nil {
		panic(err)
	}

	c, err := p.Complete(
		ali.WithText("I am the city of knowledge and Ali is its gate"),
	)
	if err != nil {
		panic(err)
	}

	text, err := c.Text()
	if err != nil {
		panic(err)
	}
	fmt.Printf("LLM says:\n%s\n", text)
}
