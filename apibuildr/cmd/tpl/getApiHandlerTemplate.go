package tpl

var GetApiHandlerTemplate = `package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/focks/apibuildr"
	"{{ .PackageName }}/internal"
	"net/http"
)

const {{ .Name }}Api = "{{ .Name }}Api"

var {{.Name}}ApiHandler = apibuildr.ApiHandler{
	Name:   {{ .Name }}Api,
	Path:   "{{ .Uri }}",
	Method: http.MethodGet,
	HandleFunc: func(w http.ResponseWriter, r *http.Request) {
		ctx := apibuildr.ApiRequestCtx(r.Context(), {{ .Name }}Api)
		w.Header().Set("request-id", apibuildr.GetRequestID(ctx))
		logger.Info(fmt.Sprintf("%s api request start", {{ .Name }}Api), apibuildr.Contextual(ctx)...)
		defer logger.Info(fmt.Sprintf("%s api request end", {{ .Name }}Api), apibuildr.Contextual(ctx)...)

		res, foul := internal.{{ .Name }}Ctrl(ctx)
		if foul != nil {
			apibuildr.HandleError(ctx, w, foul)
			return
		}
		bites, _ := json.Marshal(res)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bites)

	},
}

func init() {
	{{.Name}}ApiHandler.RegisterToRouter(Router)
}

`
