package vectorstore

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/tmc/langchaingo/schema"
)

// TODO: Actually use this lol. It populates an `id`, but is technically dead
// code. Should be used to prevent duplication based on the sha, but I think
// we're limited by pgvector/VectorStore interface current implementation for
// that. Ideally, we have an UPSERT that uses this function as a primary key.
func HashDocument(doc schema.Document) string {
	h := sha256.New()
	h.Write([]byte(doc.PageContent))

	// Sort metadata keys for consistent hashing
	keys := make([]string, 0, len(doc.Metadata))
	for k := range doc.Metadata {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		h.Write([]byte(k))
		h.Write(fmt.Appendf(nil, "%v", doc.Metadata[k]))
	}

	return hex.EncodeToString(h.Sum(nil))
}
