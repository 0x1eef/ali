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

	messages := []string{
		"Greetings.",
		"I have something important to tell you.",
		"The truth circulates with him wherever he goes.",
	}
	for _, m := range messages {
		_, err := ses.Talk(ali.WithPrompt(m))
		if err != nil {
			panic(err)
		}
	}

	if err := ses.Save("session.json"); err != nil {
		panic(err)
	}
}
