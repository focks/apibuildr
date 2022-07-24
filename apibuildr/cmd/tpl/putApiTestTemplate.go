package tpl

var PutApiHandlerTestTemplate = `package cmd
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test{{ .Name }}ApiHandler(t *testing.T) {
	loggr := getTestingLogger()

	setLogger(loggr)

	t.Run("test case", func(t *testing.T) {
		body := bytes.NewReader([]byte("{}"))
		req, err := http.NewRequest(http.MethodPut, "{{ .Uri }}", body)
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
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc({{.Name}}ApiHandler.HandleFunc)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "expecting ok")

	})

}`
