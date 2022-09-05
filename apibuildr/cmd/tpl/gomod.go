package tpl

var GoModTemplate = `module {{ .GoPackage }}

go 1.18

require github.com/stretchr/testify v1.8.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/focks/apibuildr v0.0.3 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
`
