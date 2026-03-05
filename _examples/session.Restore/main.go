package main

import (
	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
	"github.com/0x1eef/ali/session"
)

func main() {
	p, err := provider.New(ali.OpenAI)
	if err != nil {
		panic(err)
	}

	ses, err := session.New(p)
	if err != nil {
		panic(err)
	}

	if err := ses.Restore("session.json"); err != nil {
		panic(err)
	}

	c, err := ses.Talk(ali.WithText("summarize the conversation"))
	if err != nil {
		panic(err)
	}

	text, err := c.Text()
	if err != nil {
		panic(err)
	}
	println(text)
}
