package config

type EmbedderCfg struct {
	Name    string         `yaml:"name"`
	Kind    string         `yaml:"kind"`
	Options map[string]any `yaml:"options"`
}
