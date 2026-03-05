package main

import (
	"os"
	"fmt"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/image"
	"github.com/0x1eef/ali/provider"
)

func main() {
	p, err := provider.New(ali.Gemini)
	if err != nil {
		panic(err)
	}

	images, err := p.Images().Create(
		image.WithText("I am the city of knowledge and Ali is its gate"),
		image.WithQuantity(1),
	)
	if err != nil {
		panic(err)
	}

	for i, img := range images {
		f, err := os.Create(fmt.Sprintf("%d.png", i+1))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = f.ReadFrom(img)
		if err != nil {
			panic(err)
		}
	}
}
