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

func GetApiCtrlTestsTemplate() []byte {
	return []byte(`
package internal

import (
	"context"
	"fmt"
	"testing"
)

var {{ .Name }}TestCases = map[string]struct {
	description string
	wantErr     bool

	setupFunc func() error

	testValidate func(err error) error

	cleanupFunc func() error
}{
	"case 1 : base test case": {
		description: "the base case",
		wantErr:     false,
		testValidate: func(err error) error {
			return nil
		},
		setupFunc: func() error {
			fmt.Println("setup complete")
			return nil
		},
		cleanupFunc: func() error {

			return nil

		},
	},
}

func Test{{ .Name }}Ctrl(t *testing.T) {

	for testName, tt := range {{ .Name }}TestCases {
		t.Run(testName, func(t *testing.T) {
			if err := tt.setupFunc(); err != nil {
				t.Error("error setting up test for ", testName)
			}

			ctx := context.Background()
			response, err := {{ .Name }}Ctrl(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
			_ = response

			if e := tt.testValidate(err); e != nil {
				t.Error(err)
			}

			if err := tt.cleanupFunc(); err != nil {
				t.Error(err)
			}
		})
	}
}

`)
}
