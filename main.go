package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SudoBrendan/rago/cmd"
	cli "github.com/SudoBrendan/rago/pkg/cli"
	"github.com/SudoBrendan/rago/pkg/config"
	"github.com/SudoBrendan/rago/pkg/logger"
	"github.com/SudoBrendan/rago/pkg/plugins/loaders"
	"github.com/SudoBrendan/rago/pkg/plugins/models"
	"github.com/SudoBrendan/rago/pkg/plugins/vectorstores"

	// Register plugins with factories
	_ "github.com/SudoBrendan/rago/plugins/embedders"
	_ "github.com/SudoBrendan/rago/plugins/loaders"
	_ "github.com/SudoBrendan/rago/plugins/models"
	_ "github.com/SudoBrendan/rago/plugins/vectorstores"
)

func main() {
	ctx := context.Background()

	// Set up logger (dev or prod)
	log, err := logger.New(true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to set up logger: %v\n", err)
		os.Exit(1)
	}
	defer log.GetDefer()

	// Set up config
	cfg, err := config.LoadConfigFile("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}
	resolvedConfig, err := cfg.ToResolvedConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to resolve config: %v\n", err)
		os.Exit(1)
	}

	// Set up LLM Model
	llmClient, err := models.NewModelFromConfig(ctx, resolvedConfig.Model)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create LLM client: %v\n", err)
		os.Exit(1)
	}

	// Set up VectorStore
	vectorStore, err := vectorstores.NewVectorStoreFromConfig(ctx, resolvedConfig.VectorStore)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create VectorStore client: %v\n", err)
		os.Exit(1)
	}

	// Set up Loader
	loader, err := loaders.NewLoaderFromConfig(ctx, resolvedConfig.Loader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create loader: %v\n", err)
		os.Exit(1)
	}

	// Create App to hold the important stuff
	app := &cli.App{
		Model:       llmClient,
		VectorStore: vectorStore,
		Loader:      loader,
		Logger:      log,
	}

	// Run Command
	rootCmd := cmd.NewRootCmd(app)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
