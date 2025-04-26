package loaders

import (
	"context"
	"fmt"

	"github.com/SudoBrendan/rago/pkg/config"
	"github.com/tmc/langchaingo/schema"
)

// What you need to bring
type DocumentLoader interface {
	Load(ctx context.Context) ([]schema.Document, error)
	Name() string
}

type Factory func(ctx context.Context, cfg config.LoaderCfg) (DocumentLoader, error)

var registry = map[string]Factory{}

// Any plugins/loaders should Register with this in their `init()`.
func Register(kind string, factory Factory) {
	registry[kind] = factory
}

func Get(kind string) (Factory, bool) {
	f, ok := registry[kind]
	return f, ok
}

func NewLoaderFromConfig(ctx context.Context, cfg config.LoaderCfg) (DocumentLoader, error) {
	factory, ok := Get(cfg.Kind)
	if !ok {
		return nil, fmt.Errorf("loader kind %q not registered", cfg.Kind)
	}
	return factory(ctx, cfg)
}
