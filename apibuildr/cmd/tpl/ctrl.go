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

func PostApiCtrlTemplate() []byte {
	return []byte(`
package internal

import (
	"context"
	"github.com/focks/apibuildr"
	"{{ .PackageName }}/pkg/api"
)

func {{ .Name }}Ctrl(ctx context.Context, body api.{{ .Name }}ApiRequest) (*api.{{ .Name }}ApiResponse, *apibuildr.ApiFoul) {
	// todo : your business logic here
	return &api.{{ .Name }}ApiResponse{Status: "ok"}, nil
}

`)
}

func PostApiCtrlTestsTemplate() []byte {
	return []byte(`

package internal

import (
	"context"
	"errors"
	"fmt"
	"{{ .PackageName }}/pkg/api"
	"testing"
)

var {{ .Name }}TestCases = map[string]struct {
	description string
	wantErr     bool
	setupFunc func() error
	body api.{{ .Name }}ApiRequest
	testValidate func(body api.{{ .Name }}ApiRequest) error

	cleanupFunc func() error
}{
	"case 1 : base test case, everything ok": {
		description: "the base case",
		wantErr:     false,
		body: api.{{ .Name }}ApiRequest{Name: "test: rename me accordingly"},
		testValidate: func(body api.{{ .Name }}ApiRequest) error {
			res, err := {{ .Name }}Ctrl(context.Background(), body)
			
			if err != nil {
				return err 
			}
			
			if res.Status != "ok" {
				return errors.New("response status is not ok")
			}
			
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
			
			if e := tt.testValidate(tt.body); (e != nil) != tt.wantErr {
				t.Error(e)
			}

			if err := tt.cleanupFunc(); err != nil {
				t.Error(err)
			}
		})
	}
}

`)
}

func PutApiCtrlTemplate() []byte {
	return []byte(`
package internal

import (
	"context"
	"github.com/focks/apibuildr"
	"{{ .PackageName }}/pkg/api"
)

func {{ .Name }}Ctrl(ctx context.Context, body api.{{ .Name }}ApiRequest) (*api.{{ .Name }}ApiResponse, *apibuildr.ApiFoul) {
	// todo : your business logic here
	return &api.{{ .Name }}ApiResponse{Status: "ok"}, nil
}

`)
}

func PutApiCtrlTestsTemplate() []byte {
	return []byte(`
	package internal

	import (
		"context"
		"errors"
		"fmt"
		"{{ .PackageName }}/pkg/api"
		"testing"
	)
	
	var {{ .Name }}TestCases = map[string]struct {
		description string
		wantErr     bool
		setupFunc func() error
		body api.{{ .Name }}ApiRequest
		testValidate func(body api.{{ .Name }}ApiRequest) error
	
		cleanupFunc func() error
	}{
		"case 1 : base test case, everything ok": {
			description: "the base case",
			wantErr:     false,
			body: api.{{ .Name }}ApiRequest{Name: "test: rename me accordingly"},
			testValidate: func(body api.{{ .Name }}ApiRequest) error {
				res, err := {{ .Name }}Ctrl(context.Background(), body)
				
				if err != nil {
					return err 
				}
				
				if res.Status != "ok" {
					return errors.New("response status is not ok")
				}
				
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
				
				if e := tt.testValidate(tt.body); (e != nil) != tt.wantErr {
					t.Error(e)
				}
	
				if err := tt.cleanupFunc(); err != nil {
					t.Error(err)
				}
			})
		}
	}
	
	`)
}
