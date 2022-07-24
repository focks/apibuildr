package tpl

var GetApiHandlerTestTemplate = `package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	
)

func Test{{ .Name }}Handler(t *testing.T) {

	// setup 
	loggr := getTestingLogger()
	setLogger(loggr)

	t.Run("test case", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/{{.Path }}/{{.PathEnd}}", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc({{ .Name }}ApiHandler.HandleFunc)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "expecting ok")

	})
}`
