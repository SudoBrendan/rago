package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var defaultPath = filepath.Join(os.Getenv("HOME"), ".rago", "config")

/**
 *******************************************************************************
 * CONTEXTS
 *******************************************************************************
**/
type ContextCfg struct {
	Name        string `yaml:"name"`
	Model       string `yaml:"model"`
	VectorStore string `yaml:"vectorStore"`
	Loader      string `yaml:"loader"`
}

type CurrentContextCfg string

/**
 *******************************************************************************
 * CONFIGS
 *******************************************************************************
**/
type ConfigFile struct {
	Models         []ModelCfg        `yaml:"models"`
	VectorStores   []VectorStoreCfg  `yaml:"vectorStores"`
	Loaders        []LoaderCfg       `yaml:"loaders"`
	Contexts       []ContextCfg      `yaml:"contexts"`
	CurrentContext CurrentContextCfg `yaml:"current-context"`
}

// What's actually relevant for runtime
type ResolvedCurrentContextConfig struct {
	Model       ModelCfg       `yaml:"model"`
	VectorStore VectorStoreCfg `yaml:"vectorStore"`
	Loader      LoaderCfg      `yaml:"loader"`
}

func (cfg *ConfigFile) ToResolvedConfig() (*ResolvedCurrentContextConfig, error) {
	// find current-context by name
	var currentCtx *ContextCfg
	for i, ctx := range cfg.Contexts {
		if ctx.Name == string(cfg.CurrentContext) {
			currentCtx = &cfg.Contexts[i]
			break
		}
	}
	if currentCtx == nil {
		return nil, fmt.Errorf("current context %q not found in contexts", cfg.CurrentContext)
	}

	// resolve model by name
	var model *ModelCfg
	for i, m := range cfg.Models {
		if m.Name == currentCtx.Model {
			model = &cfg.Models[i]
			break
		}
	}
	if model == nil {
		return nil, fmt.Errorf("model %q not found for context %q", currentCtx.Model, currentCtx.Name)
	}

	// resolve vector store by name
	var store *VectorStoreCfg
	for i, s := range cfg.VectorStores {
		if s.Name == currentCtx.VectorStore {
			store = &cfg.VectorStores[i]
			break
		}
	}
	if store == nil {
		return nil, fmt.Errorf("vector store %q not found for context %q", currentCtx.VectorStore, currentCtx.Name)
	}

	// resolve loader by name
	var loader *LoaderCfg
	for i, l := range cfg.Loaders {
		if l.Name == currentCtx.Loader {
			loader = &cfg.Loaders[i]
			break
		}
	}
	if loader == nil {
		return nil, fmt.Errorf("loader %q not found for context %q", currentCtx.Loader, currentCtx.Name)
	}

	// return
	return &ResolvedCurrentContextConfig{
		Model:       *model,
		VectorStore: *store,
		Loader:      *loader,
	}, nil
}

/**
 *******************************************************************************
 * HELPERS
 *******************************************************************************
**/

func LoadConfigFile(path string) (*ConfigFile, error) {
	// Load file
	if path == "" {
		path = defaultPath
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %w", err)
	}
	var cfg ConfigFile
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse config: %w", err)
	}

	return &cfg, nil
}

//// Write the config file
//func SaveConfigFile(cfg *ConfigFile, path string) error {
//	if path == "" {
//		path = defaultPath
//	}
//	dir := filepath.Dir(path)
//	if err := os.MkdirAll(dir, 0755); err != nil {
//		return err
//	}
//	data, err := yaml.Marshal(cfg)
//	if err != nil {
//		return err
//	}
//	return os.WriteFile(path, data, 0644)
//}
