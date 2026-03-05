package main

import (
	"fmt"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
)

func main() {
	p, err := provider.New(ali.Gemini)
	if err != nil {
		panic(err)
	}

	c, err := p.Complete(
		ali.WithText("Describe the image"),
		ali.WithImageUrl("https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg"),
	)
	if err != nil {
		panic(err)
	}

	text, err := c.Text()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", text)
}
