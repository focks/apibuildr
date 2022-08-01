package tpl

var DeleteApiHandlerTemplate = `package cmd

import (
	"fmt"
	"net/http"

	"github.com/focks/apibuildr"
	"{{ .PackageName }}/internal"
)

const {{ .Name }}Api = "{{ .Name }}Api"

var {{ .Name }}ApiHandler = apibuildr.ApiHandler{
	Name:   {{ .Name }}Api,
	Path:   "{{ .Uri }}",
	Method: http.MethodDelete,
	HandleFunc: func(w http.ResponseWriter, r *http.Request) {
		ctx := apibuildr.ApiRequestCtx(r.Context(), {{ .Name }}Api)
		w.Header().Set("request-id", apibuildr.GetRequestID(ctx))
		logger.Info(fmt.Sprintf("%s api request start", {{ .Name }}Api), apibuildr.Contextual(ctx)...)
		defer logger.Info(fmt.Sprintf("%s api request end", {{ .Name }}Api), apibuildr.Contextual(ctx)...)

		foul := internal.{{ .Name }}Ctrl(ctx)
		if foul != nil {
			apibuildr.HandleError(ctx, w, foul)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)

	},
}

func init() {
	{{ .Name }}ApiHandler.RegisterToRouter(Router)
}
`
