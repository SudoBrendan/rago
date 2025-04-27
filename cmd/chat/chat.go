package chat

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"

	"github.com/SudoBrendan/rago/pkg/app"
)

func NewChatCmd(app *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat",
		Short: "Start a RAG-powered chat session",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Starting RAGo chat. Type 'exit' to quit.")
			var chat []llms.MessageContent
			for {
				fmt.Print("\n> ")
				prompt, _ := reader.ReadString('\n')
				prompt = strings.TrimSpace(prompt)
				app.Logger.Debug(fmt.Sprintf("Responding to %q...", prompt))

				if prompt == "exit" {
					fmt.Println("Goodbye!")
					break
				}

				finalPrompt := prompt

				// Perform retrieval if a VectorStore is configured
				var err error
				docs := []schema.Document{}
				if app.VectorStore != nil {
					docs, err = app.VectorStore.SimilaritySearch(ctx, prompt, 3, vectorstores.WithScoreThreshold(0.5))
					if err != nil {
						return fmt.Errorf("failed similarity search: %w", err)
					}

					app.Logger.Debug(fmt.Sprintf("Got the following vectorstore context: %v", docs))

					// TODO: Could add additional LLM filtering to determine relevance to the question (?)

					if len(docs) > 0 {
						contextText := "Use the following context to answer the question:\n"
						for _, doc := range docs {
							contextText += "- " + doc.PageContent + "\n"
						}
						finalPrompt = contextText + "\n" + "User Question: " + prompt
					}
				}

				// Add the message to the chat
				chat = append(chat, llms.TextParts(llms.ChatMessageTypeHuman, finalPrompt))

				// Call the model
				resp, err := app.Model.GenerateContent(ctx, chat, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
					fmt.Print(string(chunk))
					return nil
				}))
				if err != nil {
					return err
				}

				if len(docs) > 0 {
					fmt.Print("Sources: ")
					for _, d := range docs {
						if _, ok := d.Metadata["source"]; ok {
							fmt.Print(d.Metadata["source"])
						}
					}
				}

				// Save response
				for _, choice := range resp.Choices {
					chat = append(chat, llms.TextParts(llms.ChatMessageTypeAI, choice.Content))
				}
			}
			return nil
		},
	}

	return cmd
}
