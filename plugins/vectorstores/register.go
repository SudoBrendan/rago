package vectorstores

// This allows us to import `vectorstores` in main, and still get all registrations. To register a
// new `init()`, put a new import here.
import (
	_ "github.com/SudoBrendan/rago/plugins/vectorstores/pgvector"
)
