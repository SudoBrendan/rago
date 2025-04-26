package ollama

import (
	"context"

	"github.com/tmc/langchaingo/llms/ollama"

	"github.com/SudoBrendan/rago/pkg/config"
	"github.com/SudoBrendan/rago/pkg/plugins/embedders"
)

type OllamaModel struct {
	ollama.LLM
}

func NewEmbedderFromConfig(ctx context.Context, cfg config.EmbedderCfg) (embedders.Embedder, error) {
	opts, err := decodeOllamaOptions(cfg.Options)
	if err != nil {
		return nil, err
	}
	o, err := ollama.New(opts...)
	if err != nil {
		return nil, err
	}
	return &OllamaModel{
		LLM: *o,
	}, nil
}

func init() {
	embedders.Register("ollama", NewEmbedderFromConfig)
}
