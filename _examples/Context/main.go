package main

import (
	"context"
	"fmt"
	"time"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p, err := provider.New(ali.Gemini)
	if err != nil {
		panic(err)
	}

	c, err := p.Complete(
		ali.WithText("I am Ali"),
		ali.WithContext(ctx),
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
