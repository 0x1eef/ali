package assert

import (
	"testing"
)

func AssertEqual(t *testing.T, expected, actual any) {
	if expected != actual {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}

func AssertNil(t *testing.T, actual any) {
	if actual != nil {
		t.Fatalf("expected nil but got %v", actual)
	}
}
