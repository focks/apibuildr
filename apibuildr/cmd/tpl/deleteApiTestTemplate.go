package tpl

var DeleteApiHandlerTestTemplate = `package cmd
import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test{{ .Name }}ApiHandler(t *testing.T) {
	loggr := getTestingLogger()

	setLogger(loggr)

	t.Run("test case", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "{{ .Uri }}", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		
		// incase request has path parameters
		// vars := map[string]string{
		// 	"id": uuid.NewString(),
		// }
		// req = mux.SetURLVars(req, vars)
		
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc({{.Name}}ApiHandler.HandleFunc)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code, "expecting 204 no content")

	})

}`
