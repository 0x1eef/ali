## About

Ali provides a zero-dependency Go toolkit for Large Language Models that includes
OpenAI, Gemini, and Anthropic. The toolkit is young but plans full support
for chat, streaming, tool calling, audio, images, files, and structured outputs.

## Quick Start

#### Introduction

All providers implement the [ali.Provider](ali.go) interface, which serves as
the foundation for the rest of the toolkit. This ensures a consistent,
provider-agnostic API that allows implementations to be easily swapped.
The [provider.Select](provider/provider.go) function selects a provider and
automatically reads the corresponding API token from the process environment.
For example, `$OPENAI_SECRET`, `$GEMINI_SECRET`, or `$ANTHROPIC_SECRET`:

```go
package main

import (
	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
)

func main() {
	p, err := provider.Select(ali.OpenAI)
	if err != nil {
		panic(err)
	}
	// do something with 'p'
}
```

#### Complete

All providers implement a [Complete](ali.go) method that accepts a
variable number of options and returns a [ali.Completion](completion.go)
interface that is common across all providers.
See [config.go](./config.go) for a list of all available options.
The following example sends a simple prompt and prints the text
response to the terminal:


```go
package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
)

func main() {
	p, err := provider.Select(ali.OpenAI)
	if err != nil {
		panic(err)
	}

	c, err := p.Complete(
		ali.WithPrompt("Hello from #golang :)"),
	)
	if err != nil {
		panic(err)
	}

	text, err := c.Text()
	if err != nil {
		panic(err)
	}
	fmt.Printf("LLM says:\n %s\n", text)
}
```

#### Session

The [ali.Session](./session/session.go) type maintains conversation history and
context across multiple requests. Session keeps conversation state in memory and
sends prior messages on each [Talk](session/session.go) call. It is
transport/provider-neutral and works with any [ali.Provider](ali.go):

```go
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

		text, err := comp.Text()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", text)
	}
}
```

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/ali#readme)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
