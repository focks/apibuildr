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
