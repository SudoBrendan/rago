package ollama

import (
	"context"

	"github.com/SudoBrendan/rago/pkg/config"
	"github.com/SudoBrendan/rago/pkg/plugins/models"
	"github.com/tmc/langchaingo/llms/ollama"
)

type OllamaModel struct {
	ollama.LLM
}

func NewModelFromConfig(ctx context.Context, cfg config.ModelCfg) (models.Model, error) {
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
	models.Register("ollama", NewModelFromConfig)
}
