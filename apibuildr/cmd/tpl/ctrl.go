package tpl

func GetApiCtrlTemplate() []byte {
	return []byte(`
package internal

import (
	"context"
	"github.com/focks/apibuildr"
)

func {{ .Name }}Ctrl(ctx context.Context) (*apibuildr.OkResponse, *apibuildr.ApiFoul) {
	return &apibuildr.OkResponse{Status: "ok"}, nil
}

`)
}
