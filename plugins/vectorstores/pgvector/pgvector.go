package pgvector

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/vectorstores/pgvector"

	"github.com/SudoBrendan/rago/pkg/config"
	"github.com/SudoBrendan/rago/pkg/plugins/vectorstores"
)

type PgVectorStore struct {
	*pgvector.Store
}

func NewFromConfig(ctx context.Context, cfg config.VectorStoreCfg) (vectorstores.VectorStore, error) {
	// Decode options from config
	storeOpts, err := decodeExtendedPgVectorConfigOptions(ctx, cfg.Options)
	if err != nil {
		return nil, err
	}

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
