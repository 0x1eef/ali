package vcr

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/dnaeon/go-vcr.v2/cassette"
	"gopkg.in/dnaeon/go-vcr.v2/recorder"
)

// NewHTTPClient returns an HTTP client backed by a VCR recorder.
func NewHTTPClient(t *testing.T, cassetteName string) (*http.Client, func()) {
	t.Helper()

	cassettePath := filepath.Join("testdata", "cassettes", cassetteName)
	mode := recorder.ModeReplaying
	if _, err := os.Stat(cassettePath + ".yaml"); err != nil {
		if os.IsNotExist(err) {
			mode = recorder.ModeRecording
		} else {
			t.Fatalf("stat cassette: %v", err)
		}
	}

	rec, err := recorder.NewAsMode(cassettePath, mode, nil)
	if err != nil {
		t.Fatalf("new recorder: %v", err)
	}
	rec.AddSaveFilter(sanitizeInteraction)

	client := &http.Client{Transport: rec}
	stop := func() {
		if err := rec.Stop(); err != nil {
			t.Fatalf("stop recorder: %v", err)
		}
	}
	return client, stop
}

func sanitizeInteraction(i *cassette.Interaction) error {
	redactHeader(i.Request.Headers, "Authorization")
	redactHeader(i.Request.Headers, "X-Api-Key")
	redactHeader(i.Request.Headers, "x-api-key")
	u, err := url.Parse(i.Request.URL)
	if err == nil {
		q := u.Query()
		redactQuery(q, "key")
		redactQuery(q, "api_key")
		u.RawQuery = q.Encode()
		i.Request.URL = u.String()
	}
	return nil
}

func redactHeader(h http.Header, key string) {
	if _, ok := h[key]; ok {
		h.Set(key, "<redacted>")
	}
}

func redactQuery(q url.Values, key string) {
	if _, ok := q[key]; ok {
		q.Set(key, "<redacted>")
	}
}
