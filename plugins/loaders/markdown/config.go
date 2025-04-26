package markdown

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/tmc/langchaingo/textsplitter"
)

type MarkdownLoaderConfigOptions struct {
	Directory           string                    `yaml:"directory"`
	Glob                string                    `yaml:"glob"`
	TextSplitterOptions TextSplitterConfigOptions `yaml:"textSplitterOptions"`
}

type TextSplitterConfigOptions struct {
	ChunkSize    int `yaml:"chunkSize"`
	ChunkOverlap int `yaml:"chunkOverlap"`
}

func decodeMarkdownLoaderOptions(o map[string]any) (*MarkdownLoader, error) {
	var opts MarkdownLoaderConfigOptions
	if err := mapstructure.Decode(o, &opts); err != nil {
		return nil, fmt.Errorf("failed to decode markdown options: %w", err)
	}

	tsOpts := convertToTextSplitterOptions(opts.TextSplitterOptions)
	ts := textsplitter.NewMarkdownTextSplitter(tsOpts...)
	return &MarkdownLoader{
		Directory:            opts.Directory,
		Glob:                 opts.Glob,
		MarkdownTextSplitter: *ts,
	}, nil
}

func convertToTextSplitterOptions(opts TextSplitterConfigOptions) []textsplitter.Option {
	textSplitterOpts := []textsplitter.Option{}
	if opts.ChunkSize > 0 {
		textSplitterOpts = append(textSplitterOpts, textsplitter.WithChunkSize(opts.ChunkSize))
	}
	if opts.ChunkOverlap > 0 {
		textSplitterOpts = append(textSplitterOpts, textsplitter.WithChunkOverlap(opts.ChunkOverlap))
	}
	return textSplitterOpts
}
