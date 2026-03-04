package anthropic

import (
	"io"

	"github.com/0x1eef/ali"
)

type Images struct{}

func (i Images) Create(options ...func(*ali.ImageConfig)) ([]io.Reader, error) {
	return nil, ali.ErrNotImplemented
}
