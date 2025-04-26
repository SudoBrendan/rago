package embedders

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/embeddings"

	"github.com/SudoBrendan/rago/pkg/config"
)

// What you need to bring
type Embedder interface {
	embeddings.EmbedderClient

	// TODO: Custom functions
}

type Factory func(ctx context.Context, cfg config.EmbedderCfg) (Embedder, error)

var registry = map[string]Factory{}

// Plugins must register with this
func Register(kind string, factory Factory) {
	registry[kind] = factory
}

func Get(kind string) (Factory, bool) {
	f, ok := registry[kind]
	return f, ok
}

func NewEmbedderFromConfig(ctx context.Context, cfg config.EmbedderCfg) (Embedder, error) {
	factory, ok := Get(cfg.Kind)
	if !ok {
		return nil, fmt.Errorf("embedder kind %q not registered", cfg.Kind)
	}
	return factory(ctx, cfg)
}
