package vectorstore

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/tmc/langchaingo/schema"
)

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
		h.Write([]byte(fmt.Sprintf("%v", doc.Metadata[k])))
	}

	return hex.EncodeToString(h.Sum(nil))
}
