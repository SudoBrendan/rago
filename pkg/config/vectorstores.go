package config

//const (
//	VectorStoreKindPgVector = "pgvector"
//)
//
//var ValidVectorStoreKinds = map[string]struct{}{
//	VectorStoreKindPgVector: {},
//}

type VectorStoreCfg struct {
	Name     string         `yaml:"name"`
	Kind     string         `yaml:"kind"`
	Embedder EmbedderCfg    `yaml:"embedder"`
	Options  map[string]any `yaml:"options"`
}
