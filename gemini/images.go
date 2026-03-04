package gemini

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
	provider *Gemini
}

func (i Images) Create(options ...func(*ali.ImageConfig)) ([]io.Reader, error) {
	cfg := i.applyDefaults(options...)
	params := i.getParams(cfg)
	reqBody, err := i.marshal(cfg.Prompt, params)
	if err != nil {
		return nil, err
	}
	res, err := i.post(i.provider, cfg, reqBody)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return i.decode(res)
}

func (i Images) applyDefaults(options ...func(*ali.ImageConfig)) *ali.ImageConfig {
	cfg := ali.ImageConfig{Model: "imagen-4.0-generate-001", Quantity: 1}
	for _, set := range options {
		set(&cfg)
	}
	return &cfg
}

func (i Images) getParams(cfg *ali.ImageConfig) ali.Params {
	params := ali.Params{
		"sampleCount": cfg.Quantity,
	}
	for k, v := range cfg.Params {
		params[k] = v
	}
	return params
}

func (i Images) marshal(prompt string, params ali.Params) ([]byte, error) {
	type instance struct {
		Prompt string `json:"prompt"`
	}
	type payload struct {
		Parameters ali.Params `json:"parameters"`
		Instances  []instance `json:"instances"`
	}
	body := payload{
		Parameters: params,
		Instances:  []instance{{Prompt: prompt}},
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (i Images) decode(res *http.Response) ([]io.Reader, error) {
	type prediction struct {
		BytesBase64Encoded string `json:"bytesBase64Encoded"`
	}
	type payload struct {
		Predictions []prediction `json:"predictions"`
	}
	body := payload{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}
	predications := body.Predictions
	binaries := make([]io.Reader, 0, len(predications))
	for _, prediction := range predications {
		img, err := base64.StdEncoding.DecodeString(prediction.BytesBase64Encoded)
		if err != nil {
			return nil, err
		}
		binaries = append(binaries, bytes.NewReader(img))
	}
	return binaries, nil
}

func (i Images) post(g *Gemini, cfg *ali.ImageConfig, body []byte) (*http.Response, error) {
	return request.Post(
		request.WithHost(g.Host),
		request.WithPath(fmt.Sprintf("/v1beta/models/%s:predict?key=%s", cfg.Model, g.Token)),
		request.WithBody(bytes.NewReader(body)),
		request.WithClient(g.Client),
		request.WithContext(cfg.Ctx),
		request.WithSetup(func(req *http.Request) error {
			req.Header.Add("Content-Type", "application/json")
			return nil
		}),
	)
}
