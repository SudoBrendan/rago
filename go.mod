module github.com/SudoBrendan/rago

go 1.23.7

require (
	github.com/PuerkitoBio/goquery v1.8.1
	github.com/fatih/color v1.17.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/spf13/cobra v1.9.1
	github.com/stretchr/testify v1.10.0
	github.com/tmc/langchaingo v0.1.13
	go.uber.org/zap v1.27.0
	golang.org/x/text v0.22.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pgvector/pgvector-go v0.1.1 // indirect
	github.com/pkoukk/tiktoken-go v0.1.6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	gitlab.com/golang-commonmark/html v0.0.0-20191124015941-a22733972181 // indirect
	gitlab.com/golang-commonmark/linkify v0.0.0-20191026162114-a0c2df6c8f82 // indirect
	gitlab.com/golang-commonmark/markdown v0.0.0-20211110145824-bf3e522c626a // indirect
	gitlab.com/golang-commonmark/mdurl v0.0.0-20191124015652-932350d1cb84 // indirect
	gitlab.com/golang-commonmark/puny v0.0.0-20191124015043-9f83538fa04f // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.35.0 // indirect
	golang.org/x/net v0.36.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)

// TODO: update this based on direction of https://github.com/tmc/langchaingo/issues/1239
replace github.com/tmc/langchaingo => github.com/SudoBrendan/langchaingo v0.0.0-fix-pgvector-drop
