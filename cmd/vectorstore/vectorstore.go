package vectorstore

import (
	"github.com/spf13/cobra"

	"github.com/SudoBrendan/rago/pkg/app"
)

func NewVectorStoreCmd(app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vectorstore",
		Short:   "Interact with the configured vector data store.",
		Aliases: []string{"vs"},
	}
	cmd.AddCommand(NewAddDocumentsCmd(app))
	return cmd
}
