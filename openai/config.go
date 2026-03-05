package openai

// WithToken sets the OpenAI API token.
func WithToken(token string) func(o *OpenAI) {
	return func(o *OpenAI) {
		o.Token = token
	}
}

// WithHost sets the OpenAI API host.
func WithHost(host string) func(o *OpenAI) {
	return func(o *OpenAI) {
		o.Host = host
	}
}
