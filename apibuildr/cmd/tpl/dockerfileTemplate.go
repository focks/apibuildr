package tpl

var DockerfileTemplate = `# Build the manager binary
FROM golang:1.18.3 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY ./go.mod go.mod
COPY ./go.sum go.sum

RUN go mod download

# Copy the go source
COPY ./{{ .Name }} {{ .Name }}
COPY ./pkg pkg


# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bin/{{ .Name }} {{ .Name }}/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest
#FROM ubuntu:18.04
WORKDIR /
COPY --from=builder /workspace/bin/{{ .Name }} .
ENTRYPOINT ["/{{ .Name }}"]
`
