package gemini

// WithToken sets the Gemini API token.
func WithToken(token string) func(g *Gemini) {
	return func(g *Gemini) {
		g.Token = token
	}
}

// WithHost sets the Gemini API host.
func WithHost(host string) func(g *Gemini) {
	return func(g *Gemini) {
		g.Host = host
	}
}
