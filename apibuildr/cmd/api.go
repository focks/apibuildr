package cmd

import (
	"fmt"
	"github.com/focks/apibuildr/apibuildr/cmd/tpl"
	"os"
	"strings"
	"text/template"
)

type Api struct {
	Name             string
	Method           string
	Path             string
	ProjectDirectory string
}

func (a *Api) Create() error {
	// todo: check if it is a apibuildr project

	// add api file main.go
	apiFile, err := os.Create(fmt.Sprintf("%s/cmd/%s.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer apiFile.Close()

	apiTpl := tpl.GetApiTemplate()
	switch strings.ToUpper(a.Method) {
	case "GET":
	case "POST":
		apiTpl = tpl.PostApiTemplate()
		// todo: for PUT, DELETE etc
	}

	apiTemplate := template.Must(template.New("api").Parse(string(apiTpl)))
	err = apiTemplate.Execute(apiFile, a)
	if err != nil {
		return err
	}

	return nil
}
