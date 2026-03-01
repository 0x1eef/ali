package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
	"github.com/0x1eef/ali/session"
)

func main() {
	p, err := provider.Select(ali.Gemini)
	if err != nil {
		panic(err)
	}

	ses, err := session.New(p)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		prompt := scanner.Text()
		if prompt == "/exit" {
			break
		}

		comp, err := ses.Talk(ali.WithPrompt(prompt))
		if err != nil {
			panic(err)
		}

		msgs := comp.Messages()
		if len(msgs) > 0 {
			fmt.Println(msgs[len(msgs)-1].Text)
		}
	}
}
