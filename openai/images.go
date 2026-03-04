package openai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/0x1eef/ali"
	"github.com/0x1eef/ali/internal/request"
)

type Images struct {
	provider *OpenAI
}

func (i Images) Create(options ...func(*ali.ImageConfig)) ([]io.Reader, error) {
	cfg := i.applyDefaults(options...)
	params := i.build(cfg)
	reqBody, err := i.marshal(params)
	if err != nil {
		return nil, err
	}
	res, err := i.post(i.provider, reqBody, cfg)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return i.decode(res)
}

func (i Images) applyDefaults(options ...func(*ali.ImageConfig)) *ali.ImageConfig {
	cfg := ali.ImageConfig{Model: "dall-e-3", Quantity: 1}
	for _, set := range options {
		set(&cfg)
	}
	return &cfg
}

func (i Images) build(cfg *ali.ImageConfig) ali.Params {
	params := map[string]any{
		"prompt":          cfg.Prompt,
		"model":           cfg.Model,
		"n":               cfg.Quantity,
		"response_format": "b64_json",
	}
	for k, v := range cfg.Params {
		params[k] = v
	}
	return params
}

func (i Images) marshal(params ali.Params) ([]byte, error) {
	return json.Marshal(params)
}

func (i Images) decode(res *http.Response) ([]io.Reader, error) {
	type image struct {
		Base64 string `json:"b64_json"`
	}
	type payload struct {
		Data []image `json:"data"`
	}
	body := payload{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}
	images := make([]io.Reader, 0, len(body.Data))
	for _, entry := range body.Data {
		b, err := base64.StdEncoding.DecodeString(entry.Base64)
		if err != nil {
			return nil, err
		}
		images = append(images, bytes.NewReader(b))
	}
	return images, nil
}

func (i Images) post(oai *OpenAI, body []byte, cfg *ali.ImageConfig) (*http.Response, error) {
	return request.Post(
		request.WithHost(oai.Host),
		request.WithPath("/v1/images/generations"),
		request.WithBody(bytes.NewReader(body)),
		request.WithClient(oai.client),
		request.WithContext(cfg.Ctx),
		request.WithSetup(func(req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", oai.Token))
			return nil
		}),
	)
}
