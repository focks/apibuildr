package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeUriPath(t *testing.T) {

	t.Run("tc-1", func(t *testing.T) {
		uri := makeUriPath("/apibuildr/v1/hello")
		assert.Equal(t, "/apibuildr/v1/{hello:hello(?:\\\\/)?}", uri)
	})

	t.Run("tc-2", func(t *testing.T) {
		uri := makeUriPath("/apibuildr/v1/hello/{id}")
		assert.Equal(t, "/apibuildr/v1/hello/{id}", uri)
	})

	t.Run("tc-3", func(t *testing.T) {
		uri := makeUriPath("/")

		assert.Equal(t, "/", uri)
	})

	t.Run("tc-4", func(t *testing.T) {
		uri := makeUriPath("/v1")
		assert.Equal(t, "{v1:v1(?:\\\\/)?}", uri)
	})
}
