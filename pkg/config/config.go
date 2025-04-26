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

	// Validate
	// TODO: Do we even need this? Might as well just wait for the factory to load it
	// and fail then :shrug:
	//for _, m := range cfg.Models {
	//	if !isValidModelKind(m.Kind) {
	//		return nil, fmt.Errorf("unknown LLM Model Kind: '%s'", m.Kind)
	//	}
	//}
	//for _, v := range cfg.VectorStores {
	//	if !isValidVectorStoreKind(v.Kind) {
	//		return nil, fmt.Errorf("unknown LLM Model Kind: '%s'", v.Kind)
	//	}
	//}
	return &cfg, nil
}

//func isValidModelKind(kind string) bool {
//	_, ok := ValidModelKinds[kind]
//	return ok
//}

//func isValidVectorStoreKind(kind string) bool {
//	//_, ok := cv.ValidVectorStoreKinds[kind]
//	_, ok := vectorstores.Get(kind)
//	return ok
//}

// Write the config file
func SaveConfigFile(cfg *ConfigFile, path string) error {
	if path == "" {
		path = defaultPath
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

//// Given a generic model configuration, get a client for
//// that LLM
//func NewModelFromConfig(_ context.Context, cfg ModelCfg) (llms.Model, error) {
//	switch cfg.Kind {
//	case ModelKindOllama:
//		ollamaOpts, err := models.DecodeOllamaOptions(cfg.Options)
//		if err != nil {
//			return nil, err
//		}
//		return ollama.New(ollamaOpts...)
//	default:
//		return nil, fmt.Errorf("unsupported model kind: %q", cfg.Kind)
//	}
//}

//func NewVectorStoreFromConfig(ctx context.Context, cfg cv.VectorStoreCfg) (vectorstores.VectorStore, error) {
//	switch cfg.Kind {
//	case "pgvector":
//		// Create the LLM model
//		model, err := NewEmbedderFromConfig(ctx, cfg.Embedder)
//		if err != nil {
//			return pgvector.Store{}, fmt.Errorf("failed to create model for vector store %q: %w", cfg.Name, err)
//		}
//
//		// Create embedder
//		embedder, err := embeddings.NewEmbedder(model)
//		if err != nil {
//			return pgvector.Store{}, fmt.Errorf("failed to create embedder for vector store %q: %w", cfg.Name, err)
//		}
//
//		// Decode options
//		storeOpts, err := cv.DecodePGVectorOptions(cfg.Options)
//		if err != nil {
//			return pgvector.Store{}, err
//		}
//		storeOpts = append(storeOpts, pgvector.WithEmbedder(embedder))
//
//		// TODO: REMOVE
//		//storeOpts = append(storeOpts, pgvector.WithPreDeleteCollection(true))
//
//		// Create the store
//		store, err := pgvector.New(ctx, storeOpts...)
//		if err != nil {
//			return pgvector.Store{}, fmt.Errorf("failed to initialize pgvector: %w", err)
//		}
//
//		return store, nil
//	default:
//		return nil, fmt.Errorf("unsupported vector store kind: %q", cfg.Kind)
//	}
//}

//// Given a generic embedder configuration, get a client for
//// that LLM
//func NewEmbedderFromConfig(_ context.Context, cfg ModelCfg) (embeddings.EmbedderClient, error) {
//	switch cfg.Kind {
//	// TODO: ??? This introduces high coupling between llms.Model and embeddings.EmbedderClient...
//	// for the purposes of this MVP, I only use Ollama, which implements both. I imagine this could
//	// be abstracted further to distinguish between Models and Embedders... In my case, it would
//	// have just been copy-paste though :shrug: I guess the question I need to figure out is:
//	// Is an Embedder never _not_ a Model too for the sake of the YAML configuration? Probably.
//	case ModelKindOllama:
//		ollamaOpts, err := models.DecodeOllamaOptions(cfg.Options)
//		if err != nil {
//			return nil, err
//		}
//		return ollama.New(ollamaOpts...)
//	default:
//		return nil, fmt.Errorf("unsupported embedder kind: %q", cfg.Kind)
//	}
//}
