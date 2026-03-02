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
	p, err := provider.New(ali.Gemini)
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

		text, err := comp.Text()
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
	}
}
