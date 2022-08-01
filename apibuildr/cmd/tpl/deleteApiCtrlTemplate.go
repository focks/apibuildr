package tpl

var DeleteApiCtrlTemplate = `
package internal

import (
	"context"
	"github.com/focks/apibuildr"
)

func {{ .Name }}Ctrl(ctx context.Context) *apibuildr.ApiFoul {
	// todo : your business logic here
	return nil
}
`
