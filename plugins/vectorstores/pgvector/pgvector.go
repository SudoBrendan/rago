package pgvector

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/pgvector"

	"github.com/SudoBrendan/rago/pkg/config"
	"github.com/SudoBrendan/rago/pkg/plugins/embedders"
	"github.com/SudoBrendan/rago/pkg/plugins/vectorstores"
)

type PgVectorStore struct {
	*pgvector.Store
}

func NewFromConfig(ctx context.Context, cfg config.VectorStoreCfg) (vectorstores.VectorStore, error) {
	// Create the LLM model
	model, err := embedders.NewEmbedderFromConfig(ctx, cfg.Embedder)
	if err != nil {
		return nil, fmt.Errorf("failed to create model for vector store %q: %w", cfg.Name, err)
	}

	// Create embedder
	embedder, err := embeddings.NewEmbedder(model)
	if err != nil {
		return nil, fmt.Errorf("failed to create embedder for vector store %q: %w", cfg.Name, err)
	}

	// Decode options from config
	storeOpts, err := decodeExtendedPgVectorConfigOptions(cfg.Options)
	if err != nil {
		return nil, err
	}

	storeOpts = append(storeOpts, pgvector.WithEmbedder(embedder))

	// Create the store
	store, err := pgvector.New(ctx, storeOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize pgvector: %w", err)
	}

	return PgVectorStore{
		Store: &store,
	}, nil
}

func init() {
	vectorstores.Register("pgvector", NewFromConfig)
}
