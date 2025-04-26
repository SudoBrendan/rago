/**Required for PGVector support
create extension if not exists vector;

/*Required for UPSERT support*/
CREATE TABLE IF NOT EXISTS your_table (
  id TEXT PRIMARY KEY,
  content TEXT,
  embedding VECTOR,
  metadata JSONB
);

**/