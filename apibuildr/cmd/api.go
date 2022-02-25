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
	PackageName      string
}

func (a *Api) Create() error {

	_, err := os.Stat(fmt.Sprintf("%s/main.go", a.ProjectDirectory))
	if err != nil {
		CheckError("the current directory is not apibuildr project.")
	}
	_, err = os.Stat(fmt.Sprintf("%s/cmd/server.go", a.ProjectDirectory))
	if err != nil {
		CheckError("the current directory is not apibuildr project.")
	}
	_, err = os.Stat(fmt.Sprintf("%s/cmd/%s.go", a.ProjectDirectory, a.Name))
	if err == nil {
		CheckError("the given api name already exists.")
	}

	// create the internal directory
	internalMod := fmt.Sprintf("%s/internal", a.ProjectDirectory)
	if _, err := os.Stat(internalMod); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(internalMod, 0754); err != nil {
			return err
		}
	}

	switch strings.ToUpper(a.Method) {
	case "GET":
		return a.createGetApi()
	case "POST":

	}

	// checking if internal directory exists

	// add ctrl file
	ctrlFile, err := os.Create(fmt.Sprintf("%s/internal/%s.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer ctrlFile.Close()
	ctrlTpl := tpl.GetApiCtrlTemplate()
	ctrlTemplate := template.Must(template.New("ctrl").Parse(string(ctrlTpl)))
	err = ctrlTemplate.Execute(ctrlFile, a)
	if err != nil {
		return err
	}

	return nil
}

func (a *Api) createGetApi() error {
	// add api file
	apiFile, err := os.Create(fmt.Sprintf("%s/cmd/%s.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer apiFile.Close()
	apiTpl := tpl.GetApiTemplate()
	apiTemplate := template.Must(template.New("api").Parse(string(apiTpl)))
	err = apiTemplate.Execute(apiFile, a)
	if err != nil {
		return err
	}

	// checking if internal directory exists
	internalMod := fmt.Sprintf("%s/internal", a.ProjectDirectory)
	if _, err := os.Stat(internalMod); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(internalMod, 0754); err != nil {
			return err
		}
	}
	// add ctrl file
	ctrlFile, err := os.Create(fmt.Sprintf("%s/internal/%s.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer ctrlFile.Close()
	ctrlTpl := tpl.GetApiCtrlTemplate()
	ctrlTemplate := template.Must(template.New("ctrl").Parse(string(ctrlTpl)))
	err = ctrlTemplate.Execute(ctrlFile, a)
	if err != nil {
		return err
	}

	return nil
}
