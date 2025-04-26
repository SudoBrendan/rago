package rag

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/vectorstores"
)

type RAGClient struct {
	store vectorstores.VectorStore
	llm   llms.Model
}

// RAGClient implements llms.Model and injects RAG context before responding.
func NewRAGClient(ctx context.Context, llm llms.Model, ragStore vectorstores.VectorStore) *RAGClient {
	return &RAGClient{
		store: ragStore,
		llm:   llm,
	}
}

// Call proxies directly to the underlying model (text-only mode, no RAG).
//
// Deprecated: Call is retained for compatibility. Use GenerateContent instead.
func (r *RAGClient) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	return llms.GenerateFromSinglePrompt(ctx, r, prompt, options...)
}

// Inject RAG before calling downstream LLM.
func (r *RAGClient) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error) {
	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages provided")
	}

	// 1. Get the most recent message and use it as our query
	var query string

outer:
	for i := len(messages) - 1; i >= 0; i-- {
		msg := messages[i]
		if msg.Role != llms.ChatMessageTypeSystem {
			for _, part := range msg.Parts {
				if textPart, ok := part.(llms.TextContent); ok {
					// only use the first TextContent we find
					query = textPart.String()
					break outer
				}
			}
			// only attempt to look at the first non-system message
			break
		}
	}

	// 2. Perform RAG retrieval if we have a message
	if query != "" {
		// TODO: make k and additional options configurable
		docs, err := r.store.SimilaritySearch(ctx, query, 3)
		if err != nil {
			return nil, fmt.Errorf("rag retrieval failed: %w", err)
		}

		// 3. Construct a system message with context
		var contextText string
		for _, doc := range docs {
			contextText += doc.PageContent + "\n"
		}

		systemMessage := llms.MessageContent{
			Role: llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{
				llms.TextContent{Text: "Use the following context to answer the question:\n" + contextText},
			},
		}

		// 4. Prepend system context and pass full message list to LLM
		augmentedMessages := append([]llms.MessageContent{systemMessage}, messages...)
		return r.llm.GenerateContent(ctx, augmentedMessages, options...)
	}

	// fallback to no RAG
	return r.llm.GenerateContent(ctx, messages, options...)
}
