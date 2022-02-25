package tpl

func GetApiTemplate() []byte {
	return []byte(`
package cmd

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
	Path:   "{{ .Path }}",
	Method: http.MethodGet,
	HandleFunc: func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		defer fmt.Println("api request ends")
		fmt.Println("api request starts")
		ctx := apibuildr.ApiRequestCtx(r.Context(), {{ .Name }}Api)

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

`)
}

func PostApiTemplate() []byte {
	return []byte(`
package cmd

import (
	"fmt"
	"net/http"
)

const {{ .Name }}Api = "{{ .Name }}Api"

var helloApiHandler = ApiHandler{
	Name:   HelloApi,
	Path:   "/v1/{{{ .Path }}:{{ .Path }}(?:\\/)?}",
	Method: http.MethodGet,
	HandleFunc: func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		defer fmt.Println("api request ends")
		fmt.Println("api request starts")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte())

	},
}

func init() {
	helloApiHandler.RegisterToRouter(Router)
}

`)
}
