package ollama

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/tmc/langchaingo/llms/ollama"
)

type OllamaOptions struct {
	Model     string `yaml:"model"`
	ServerURL string `yaml:"serverURL"`
}

func decodeOllamaOptions(o map[string]any) ([]ollama.Option, error) {
	var opts OllamaOptions
	if err := mapstructure.Decode(o, &opts); err != nil {
		return nil, fmt.Errorf("failed to decode Ollama options: %w", err)
	}

	ollamaOpts := []ollama.Option{}
	if opts.Model != "" {
		ollamaOpts = append(ollamaOpts, ollama.WithModel(opts.Model))
	}
	if opts.ServerURL != "" {
		ollamaOpts = append(ollamaOpts, ollama.WithServerURL(opts.ServerURL))
	}
	return ollamaOpts, nil
}
