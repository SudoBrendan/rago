package cmd

import (
	"github.com/spf13/cobra"

	"github.com/SudoBrendan/rago/cmd/chat"
	"github.com/SudoBrendan/rago/cmd/vectorstore"
	"github.com/SudoBrendan/rago/pkg/app"
)

func NewRootCmd(app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rago",
		Short: "RAGo is a golang CLI implementation of configurable RAG",
		Long:  `RAGo is a CLI that allows you to work with pluggable RAG implementations and iterate on them using kubectl-like configuration.`,
	}

	cmd.AddCommand(vectorstore.NewVectorStoreCmd(app))
	cmd.AddCommand(chat.NewChatCmd(app))

	return cmd
}
