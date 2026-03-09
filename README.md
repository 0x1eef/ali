## About

Ali is designed as a minimal, composable foundation for building LLM-powered
applications in Go. It focuses on small interfaces, explicit configuration,
and compatibility across providers. Ali has zero dependencies outside Go's
standard library.

## Quick Start

#### session.Talk

[ali.Session](./session/session.go) maintains conversation history and
context across multiple requests. A session stores conversation history
in memory and automatically includes prior messages on each [Talk](session/session.go)
call. It is transport/provider-neutral and works with any [ali.Provider](ali.go).
The following example implements a simple **R**ead **E**val **P**rint **L**oop
with the help of the [ali.Session](./session/session.go):

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

		c, err := ses.Talk(ali.WithText(prompt))
		if err != nil {
			panic(err)
		}

		text, err := c.Text()
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
	}
}
```

#### session.{Save,Restore}

A session can be saved to disk and afterwards restored via the
[Session.Save](session/session.go) and
[Session.Restore](session/session.go)
methods. Going further &ndash;
a session can be written to any io.Writer and read from any io.Reader
via the [session.WriteTo](session/session.go) and [session.ReadFrom](session/session.go)
methods.

This opens the door for more than just writing to or reading from files on disk,
and creates possibilities like storing a session in a database of some kind &ndash;
for example, a JSONB column in a PostgreSQL database would be perfect:

```go
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
		_, err := ses.Talk(ali.WithText(m))
		if err != nil {
			panic(err)
		}
	}

	if err := ses.Save("session.json"); err != nil {
		panic(err)
	}
}
```

#### Pool

Ali creates a new [http.Client](https://pkg.go.dev/net/http#Client) for each
request by default. If you want finer transport control, use
[ali.WithClient](config.go) to provide your own client. The example below uses
a custom [http.Transport](https://pkg.go.dev/net/http#Transport) to configure
a connection pool and applies it once at session construction, so all
[session.Talk](session/session.go) calls reuse the same connection pool:

```go
package main

import (
	"net/http"
	"time"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
	"github.com/0x1eef/ali/session"
)

func main() {
	p, err := provider.New(ali.OpenAI)
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		MaxConnsPerHost:     50,
		IdleConnTimeout:     90 * time.Second,
	}
	pool := &http.Client{Transport: transport}

	ses, err := session.New(p, ali.WithClient(pool))
	if err != nil {
		panic(err)
	}

	messages := []string{
		"Explain connection pooling in one sentence.",
		"Now explain it with an analogy.",
	}
	for _, m := range messages {
		_, err = ses.Talk(ali.WithText(m))
		if err != nil {
			panic(err)
		}
	}
}
```

## Features

#### Architecture

* 🧩 Small, focused interfaces
* 🔄 Provider-agnostic abstractions
* ⚙️ Explicit configuration
* 🚫 No global state

#### Providers

* 🌐 OpenAI, Gemini, and Anthropic providers
* 🔌 Automatic environment-based token loading via [provider.New](provider/provider.go)
* 🧱 Direct provider constructors via [openai.New](openai/openai.go) and friends
* 🛰️ Support for providers with an OpenAI-compatible API via [openai.WithHost](openai/config.go)

#### Requests

* 🗂️ Stateless one-shot completions via [ali.Provider.Complete](ali.go)
* 🛠️ Composable request options via [ali.WithText](config.go), [ali.WithRole](config.go) and friends
* 🌊 Connection pool support via [ali.WithClient](config.go) and [http.Transport](https://pkg.go.dev/net/http#Transport)
* 🧠 Multimodal inputs via [ali.WithText](config.go), [ali.WithImageUrl](config.go), and [ali.WithPdf](config.go)
* 🖼️ Image generation via [ali.Provider.Images](ali.go)

#### Sessions

* 💬 In-memory multi-turn conversations via [session.Session](session/session.go)
* 🔁 Conversation continuity via [session.Talk(...)](session/session.go)
* 💾 Session persistence via [session.Save](session/session.go), [session.Restore](session/session.go), and friends

#### Completions

* 📊 Unified completion access (`Text`, `InputTokens`, `OutputTokens`, `TotalTokens`)
* 🔎 Raw provider response access with `Raw()`

#### Dependencies

* 📦 Zero dependencies outside Go standard library


## Examples

#### ali.Provider

All providers implement the [ali.Provider](ali.go) interface, which serves as
the foundation for the rest of the toolkit. This ensures a consistent,
provider-agnostic API that allows implementations to be easily swapped.
The [provider.New](provider/provider.go) function returns a type that implements
the [ali.Provider](ali.go) interface and automatically reads the corresponding
API token from the process environment.

For example &ndash; `$OPENAI_SECRET`, `$GEMINI_SECRET`, or `$ANTHROPIC_SECRET`:

```go
package main

import (
	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/provider"
)

func main() {
	p, err := provider.New(ali.OpenAI)
	if err != nil {
		panic(err)
	}
	// do something with 'p'
}
```

But sometimes explicit configuration is preferred &ndash; for example, when a
caller wants to import only OpenAI-specific code. Otherwise &ndash; when a caller
imports the [provider](provider/provider.go) package they also import the OpenAI,
Anthropic, and Gemini packages as well.

In that scenario &ndash; and others like it &ndash; a provider can be built
directly instead of using [provider.New](provider/provider.go). The example below
uses OpenAI, but the same approach works for Anthropic and Gemini. This approach
is a little more verbose, which is why [provider.New](provider/provider.go)
exists in the first place but remains totally valid:

```go
package main

import (
	"github.com/0x1eef/ali/openai"
)

func main() {
	p, err := openai.New(
		openai.WithToken("yourtoken"),
	)
	if err != nil {
		panic(err)
	}
	// do something with 'p'
}
```

#### Complete

All providers implement a [Complete](ali.go) method that accepts a
variable number of options and returns a [ali.Completion](ali.go)
interface that is common across all providers. This method is stateless
and does not carry state between method calls.  See [config.go](./config.go)
for a list of all available options. The following example sends a simple
prompt and prints the text response to the terminal:

```go
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
```

#### Context

Every kind of request that Ali makes can be covered by the [context](https://pkg.go.dev/context)
package, and this can give the caller greater control over the requests
that Ali makes. For example, and perhaps most common, a context can be used
to implement a request timeout that results in an error when the limit
naturally expires:

```go
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
```

#### Images

Both OpenAI and Gemini implement the [ali.ImageProvider](ali.go) interface
but Anthropic does not. The following example uses a type assertion because
not every provider implements the [ali.ImageProvider](ali.go) interface, and
the type assertion acts as a guard.

But it is possible to avoid the type assertion when you instantiate a provider
directly (eg [openai.New](openai/openai.go)) instead of using [provider.New](provider/provider.go).
The following example uses Gemini, and the [imagen](https://ai.google.dev/gemini-api/docs/imagen)
model to generate an image from a prompt that is then written to disk:

```go
package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/image"
	"github.com/0x1eef/ali/provider"
)

func main() {
	p, err := provider.New(ali.Gemini)
	if err != nil {
		panic(err)
	}

	imgp, ok := p.(ali.ImageProvider)
	if !ok {
		fmt.Printf("%s does not support image generation\n", p.Name())
		os.Exit(1)
	}

	images, err := imgp.Images().Create(
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
```

#### Multimodal

[ali.WithText](config.go) sends text input, but a request does not need to be
text-only. Providers can accept multimodal input where a single message has
multiple parts. In Ali, this is done with options like [ali.WithText](config.go),
[ali.WithPdf](config.go), [ali.WithImageUrl](config.go), and friends &ndash; all
of which can be used together in the same request, or independently of each other.

For example, a common scenario is to have one part that asks a question as
text, and another part that is the subject of the question. The subject might
be a local file, or the URL to an image &ndash; both of which are treated as
non-text inputs by providers and represented differently in a request:

```go
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
		ali.WithText("Describe the image"),
		ali.WithImageUrl("https://upload.wikimedia.org/wikipedia/commons/3/3f/Fronalpstock_big.jpg"),
		ali.WithText("Summarize the book"),
		ali.WithPdf("book.pdf"),
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
```

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/ali#readme)
* [codeberg.org/@0x1eef](https://codeberg.org/0x1eef/ali)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
