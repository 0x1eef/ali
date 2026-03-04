package image

import (
	"github.com/0x1eef/ali"
)

func WithPrompt(prompt string) func(*ali.ImageConfig) {
	return func(r *ali.ImageConfig) {
		r.Prompt = prompt
	}
}

func WithQuantity(quantity int) func(*ali.ImageConfig) {
	return func(r *ali.ImageConfig) {
		r.Quantity = quantity
	}
}

func WithModel(model string) func(*ali.ImageConfig) {
	return func(r *ali.ImageConfig) {
		r.Model = model
	}
}

func WithParams(params ali.Params) func(*ali.ImageConfig) {
	return func(r *ali.ImageConfig) {
		r.Params = params
	}
}
