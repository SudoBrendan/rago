package pgvector

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

type ExtendedPgVectorConfigOptions struct {
	ConnectionURL       string `yaml:"connectionURL"`
	PreDeleteCollection bool   `yaml:"preDeleteCollection"`
}

func decodeExtendedPgVectorConfigOptions(o map[string]any) ([]pgvector.Option, error) {
	var opts ExtendedPgVectorConfigOptions
	if err := mapstructure.Decode(o, &opts); err != nil {
		return nil, fmt.Errorf("failed to decode pgvector options: %w", err)
	}

	pgvectorOpts := []pgvector.Option{}
	if opts.ConnectionURL != "" {
		pgvectorOpts = append(pgvectorOpts, pgvector.WithConnectionURL(opts.ConnectionURL))
	}
	if opts.PreDeleteCollection {
		pgvectorOpts = append(pgvectorOpts, pgvector.WithPreDeleteCollection(opts.PreDeleteCollection))
	}
	return pgvectorOpts, nil
}
