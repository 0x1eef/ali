package anthropic

// WithToken sets the Anthropic API token.
func WithToken(token string) func(o *Anthropic) {
	return func(a *Anthropic) {
		a.token = token
	}
}

// WithHost sets the Anthropic API host.
func WithHost(host string) func(o *Anthropic) {
	return func(a *Anthropic) {
		a.host = host
	}
}
