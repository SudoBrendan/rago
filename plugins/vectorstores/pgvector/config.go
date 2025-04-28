package pgvector

import (
	"context"
	"fmt"

	"github.com/SudoBrendan/rago/pkg/config"
	"github.com/SudoBrendan/rago/pkg/plugins/embedders"
	"github.com/mitchellh/mapstructure"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

type ExtendedPgVectorConfigOptions struct {
	ConnectionURL       string              `yaml:"connectionURL"`
	PreDeleteCollection bool                `yaml:"preDeleteCollection"`
	Embedder            *config.EmbedderCfg `yaml:"embedder"`
}

func decodeExtendedPgVectorConfigOptions(ctx context.Context, o map[string]any) ([]pgvector.Option, error) {
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
	if opts.Embedder != nil {
		// Create the LLM model
		model, err := embedders.NewEmbedderFromConfig(ctx, *opts.Embedder)
		if err != nil {
			return nil, fmt.Errorf("failed to create embedder model for vector store: %w", err)
		}

		// Create embedder
		embedder, err := embeddings.NewEmbedder(model)
		if err != nil {
			return nil, fmt.Errorf("failed to create embedder for vector store: %w", err)
		}
		pgvectorOpts = append(pgvectorOpts, pgvector.WithEmbedder(embedder))
	}

	return pgvectorOpts, nil
}
