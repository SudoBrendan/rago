package cmd

import (
	"github.com/tmc/langchaingo/llms"

	"github.com/SudoBrendan/rago/pkg/logger"
	"github.com/SudoBrendan/rago/pkg/plugins/loaders"
	"github.com/SudoBrendan/rago/pkg/plugins/vectorstores"
)

type App struct {
	Model       llms.Model
	VectorStore vectorstores.VectorStore
	Loader      loaders.DocumentLoader
	Logger      logger.Logger
}
