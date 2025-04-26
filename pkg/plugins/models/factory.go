package models

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"

	"github.com/SudoBrendan/rago/pkg/config"
)

// What you need to bring
type Model interface {
	llms.Model

	// TODO: Custom functions
}

type Factory func(ctx context.Context, cfg config.ModelCfg) (Model, error)

var registry = map[string]Factory{}

// Plugins must register with this
func Register(kind string, factory Factory) {
	registry[kind] = factory
}

func Get(kind string) (Factory, bool) {
	f, ok := registry[kind]
	return f, ok
}

func NewModelFromConfig(ctx context.Context, cfg config.ModelCfg) (Model, error) {
	factory, ok := Get(cfg.Kind)
	if !ok {
		return nil, fmt.Errorf("model kind %q not registered", cfg.Kind)
	}
	return factory(ctx, cfg)
}
