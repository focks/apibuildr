package tpl

var MakefileTemplate = `
IMG ?= {{ .Name }}:0.0.1

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif


# Run tests
test: fmt vet
	go test {{ .PackageName }}/{{ .Name }}/cmd/tests -v -coverprofile {{ .Name }}-cover.out

build: fmt vet
	go build -o bin/{{ .Name }} ./{{ .Name }}/main.go

run: fmt vet
	go run ./{{ .Name }}/main.go

fmt:
	go fmt ./{{ .Name }}/...

vet:
	go vet ./{{ .Name }}/...

# Build the docker image
docker-build:
	docker build . -t ${IMG} -f ./{{ .Name }}/Dockerfile

# Push the docker image
docker-push:
	docker push ${IMG}

deploy:
	docker push ${IMG}
`
