package vectorstore

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SudoBrendan/rago/pkg/app"
)

func NewAddDocumentsCmd(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:     "add-documents",
		Short:   "Add documents to the configured vector store using the configured loader",
		Aliases: []string{"add"},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			// Load the documents
			docs, err := app.Loader.Load(ctx)
			if err != nil {
				return fmt.Errorf("failed to load documents: %w", err)
			}

			// Generate distinct IDs for each document so we don't duplicate records
			for i := range docs {
				id := HashDocument(docs[i])
				if docs[i].Metadata == nil {
					docs[i].Metadata = map[string]any{}
				}
				docs[i].Metadata["id"] = id
			}

			// Write to the vector store
			if _, err := app.VectorStore.AddDocuments(ctx, docs); err != nil {
				return fmt.Errorf("failed to add documents to vector store: %w", err)
			}

			app.Logger.Info(fmt.Sprintf("✅ Added %d documents to vector store\n", len(docs)))
			return nil
		},
	}
}
