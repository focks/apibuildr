package tpl

var PutApiHandlerTemplate = `package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/focks/apibuildr"
	"{{ .PackageName }}/internal"
	"{{ .PackageName }}/pkg/api"
	"io/ioutil"
	"net/http"
)

const {{ .Name }}Api = "{{ .Name }}Api"

var {{ .Name }}ApiHandler = apibuildr.ApiHandler{
	Name:   {{ .Name }}Api,
	Path:   "/{{ .Uri }}",
	Method: http.MethodPut,
	HandleFunc: func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		
		ctx := apibuildr.ApiRequestCtx(r.Context(), {{ .Name }}Api)
		w.Header().Set("request-id", apibuildr.GetRequestID(ctx))
		logger.Info(fmt.Sprintf("%s api request start", {{ .Name }}Api), apibuildr.Contextual(ctx)...)
		defer logger.Info(fmt.Sprintf("%s api request end", {{ .Name }}Api), apibuildr.Contextual(ctx)...)

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			apibuildr.HandleError(ctx, w, err)
			return
		}
		var request api.{{ .Name }}ApiRequest

		if err = json.Unmarshal(bodyBytes, &request); err != nil {
			apibuildr.HandleError(ctx, w, err)
			return
		}

		res, foul := internal.{{ .Name }}Ctrl(ctx, request)
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
	{{ .Name }}ApiHandler.RegisterToRouter(Router)
}
`
