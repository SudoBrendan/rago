package vectorstore

import (
	"github.com/spf13/cobra"

	app "github.com/SudoBrendan/rago/pkg/cli"
)

func NewVectorStoreCmd(app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vectorstore",
		Short:   "Interact with the configured vector store",
		Aliases: []string{"vs"},
	}
	cmd.AddCommand(NewAddDocumentsCmd(app))
	return cmd
}
