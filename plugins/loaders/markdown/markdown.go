package markdown

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"

	"github.com/SudoBrendan/rago/pkg/config"
	"github.com/SudoBrendan/rago/pkg/plugins/loaders"
)

// MarkdownLoader implements DocumentLoader interface.
type MarkdownLoader struct {
	textsplitter.MarkdownTextSplitter

	// Custom
	Directory string
	Glob      string
}

func NewFromConfig(ctx context.Context, cfg config.LoaderCfg) (loaders.DocumentLoader, error) {
	// Create the Markdown textsplitter
	loader, err := decodeMarkdownLoaderOptions(cfg.Options)
	if err != nil {
		return nil, err
	}

	// TODO: write a better structure for required and default values, likely in pkg/config.
	if loader.Directory == "" {
		return nil, fmt.Errorf("no directory found in markdown loader config")
	}
	if loader.Glob == "" {
		loader.Glob = "*.md"
	}
	return loader, nil
}

func (m MarkdownLoader) Load(ctx context.Context) ([]schema.Document, error) {
	var files []string
	err := filepath.WalkDir(m.Directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		match, _ := filepath.Match(m.Glob, filepath.Base(path))
		if !d.IsDir() && match {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}

	var docs []schema.Document
	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		chunks, err := m.SplitText(string(content))
		if err != nil {
			return nil, err
		}
		for _, c := range chunks {
			docs = append(docs, schema.Document{
				PageContent: c,
				Metadata: map[string]any{
					"source": f,
				},
			})
		}
	}

	return docs, nil
}

func init() {
	loaders.Register("markdown", NewFromConfig)
}
