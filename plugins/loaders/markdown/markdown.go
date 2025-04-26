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
	//mdOpts, err := decodeTextSplitterOptions(cfg.Options)
	//if err != nil {
	//	return nil, err
	//}
	//ts := textsplitter.NewMarkdownTextSplitter(mdOpts...)

	// TODO: write a better structure for required and default values, likely in pkg/config.
	if loader.Directory == "" {
		return nil, fmt.Errorf("no directory found in markdown loader config")
	}
	if loader.Glob == "" {
		loader.Glob = "*.md"
	}
	return loader, nil

	//return MarkdownLoader{
	//	Directory:            cfg.Directory,
	//	Glob:                 cfg.Glob,
	//	MarkdownTextSplitter: *ts,
	//}, nil
}

//func New(options map[string]any) (*MarkdownLoader, error) {
//	dir, _ := options["directory"].(string)
//	if dir == "" {
//		return nil, fmt.Errorf("missing required option: directory")
//	}
//	glob, _ := options["glob"].(string)
//	if glob == "" {
//		glob = "*.md"
//	}
//
//	// Optional splitter config
//	chunkSize := 300
//	chunkOverlap := 50
//	if val, ok := options["chunkSize"].(int); ok {
//		chunkSize = val
//	}
//	if val, ok := options["chunkOverlap"].(int); ok {
//		chunkOverlap = val
//	}
//
//	splitter := textsplitter.NewRecursiveCharacter(
//		textsplitter.WithChunkSize(chunkSize),
//		textsplitter.WithChunkOverlap(chunkOverlap),
//	)
//
//	return &MarkdownLoader{
//		Dir:      dir,
//		Glob:     glob,
//		Splitter: splitter,
//	}, nil
//}

func (m MarkdownLoader) Name() string {
	return "markdown"
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
	// Register this loader with our factory as this unique name
	loaders.Register("markdown", NewFromConfig)
}
