package vectorstores

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/vectorstores"

	"github.com/SudoBrendan/rago/pkg/config"
)

// What you need to bring
type VectorStore interface {
	vectorstores.VectorStore

	// TODO: Custom functions
}

type Factory func(ctx context.Context, cfg config.VectorStoreCfg) (VectorStore, error)

var registry = map[string]Factory{}

// Plugins must register with this
func Register(kind string, factory Factory) {
	registry[kind] = factory
}

func Get(kind string) (Factory, bool) {
	f, ok := registry[kind]
	return f, ok
}

func NewVectorStoreFromConfig(ctx context.Context, cfg config.VectorStoreCfg) (VectorStore, error) {
	factory, ok := Get(cfg.Kind)
	if !ok {
		return nil, fmt.Errorf("vector store kind %q not registered", cfg.Kind)
	}
	return factory(ctx, cfg)
}
