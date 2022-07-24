package tpl

var PostApiCtrlTemplate = `package internal

import (
	"context"
	"github.com/focks/apibuildr"
	"{{ .PackageName }}/pkg/api"
)

func {{ .Name }}Ctrl(ctx context.Context, body api.{{ .Name }}ApiRequest) (*api.{{ .Name }}ApiResponse, *apibuildr.ApiFoul) {
	// todo : your business logic here
	return &api.{{ .Name }}ApiResponse{Status: "ok"}, nil
}

`
