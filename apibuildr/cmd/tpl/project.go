package tpl

func MainTemplate() []byte {
	return []byte(`
package main

import "{{ .PackageName }}/cmd"

func main() {
	cmd.Exec()
}
`)
}

func InitFileTemplate() []byte {
	return []byte(`package {{ .Package }}

import "go.uber.org/zap"

var logger *zap.Logger

func Initialize(lg *zap.Logger) {
	logger = lg 
}
`)
}
