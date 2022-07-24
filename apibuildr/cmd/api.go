package cmd

import (
	"fmt"

	"os"
	"strings"

	"github.com/focks/apibuildr/apibuildr/cmd/tpl"
)

type Api struct {
	Name             string
	Method           string
	Path             string
	PathEnd          string
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
	if err := createInternalDirectory(a.ProjectDirectory); err != nil {
		CheckError(err)
	}

	if _, err := os.Stat(fmt.Sprintf("%s/internal/init.go", a.ProjectDirectory)); os.IsNotExist(err) {
		if err := createInitFile(InitFileVars{
			Package: "internal",
			Path:    fmt.Sprintf("%s/internal", a.ProjectDirectory),
		}); err != nil {
			return err
		}
	}

	switch strings.ToUpper(a.Method) {
	case "GET":
		return a.createGetApi()
	case "POST":
		return a.createPostApi()
	case "PUT":
		return a.createPutApi()

	default:
		return nil

	}
}

func createInternalDirectory(projectDir string) error {
	// create the internal directory
	internalMod := fmt.Sprintf("%s/internal", projectDir)
	if _, err := os.Stat(internalMod); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(internalMod, 0754); err != nil {
			return err
		}
	}

	return nil
}

func (a *Api) createGetApi() error {
	files := []*ApiFile{
		{
			path:     fmt.Sprintf("%s/cmd/%sHandler.go", a.ProjectDirectory, a.Name),
			template: tpl.GetApiHandlerTemplate,
		},
		{
			path:     fmt.Sprintf("%s/cmd/%sHandler_test.go", a.ProjectDirectory, a.Name),
			template: tpl.GetApiHandlerTestTemplate,
		},
		{
			path:     fmt.Sprintf("%s/internal/%sCtrl.go", a.ProjectDirectory, a.Name),
			template: tpl.GetApiCtrlTemplate,
		},
	}

	if err := a.realizeApiDirectories(); err != nil {
		return err
	}

	a.makeFiles(&files)

	return nil
}

func (a *Api) createPostApi() error {
	files := []*ApiFile{
		{
			path:     fmt.Sprintf("%s/cmd/%sHandler.go", a.ProjectDirectory, a.Name),
			template: tpl.PostApiHandlerTemplate,
		},
		{
			path:     fmt.Sprintf("%s/cmd/%sHandler_test.go", a.ProjectDirectory, a.Name),
			template: tpl.PostApiHandlerTestTemplate,
		},
		{
			path:     fmt.Sprintf("%s/internal/%sCtrl.go", a.ProjectDirectory, a.Name),
			template: tpl.PostApiCtrlTemplate,
		},
		{
			path:     fmt.Sprintf("%s/pkg/api/%sReqRes.go", a.ProjectDirectory, a.Name),
			template: tpl.RequestResponseTemplate,
		},
	}

	if err := a.realizeApiDirectories(); err != nil {
		return err
	}

	a.makeFiles(&files)

	return nil
}

func (a *Api) createPutApi() error {
	files := []*ApiFile{
		{
			path:     fmt.Sprintf("%s/cmd/%sHandler.go", a.ProjectDirectory, a.Name),
			template: tpl.PutApiHandlerTemplate,
		},
		// {
		// 	path:     fmt.Sprintf("%s/cmd/%sHandler_test.go", a.ProjectDirectory, a.Name),
		// 	template: tpl.PutApiN(),
		// },
		{
			path:     fmt.Sprintf("%s/internal/%sCtrl.go", a.ProjectDirectory, a.Name),
			template: tpl.PutApiCtrlTemplate,
		},
		{
			path:     fmt.Sprintf("%s/pkg/api/%sReqRes.go", a.ProjectDirectory, a.Name),
			template: tpl.RequestResponseTemplate,
		},
	}

	if err := a.realizeApiDirectories(); err != nil {
		return err
	}

	a.makeFiles(&files)

	return nil
}
