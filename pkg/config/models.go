package config

type ModelCfg struct {
	Name    string         `yaml:"name"`
	Kind    string         `yaml:"kind"`
	Options map[string]any `yaml:"options"`
}
