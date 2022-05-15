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
	internalMod := fmt.Sprintf("%s/internal", a.ProjectDirectory)
	if _, err := os.Stat(internalMod); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(internalMod, 0754); err != nil {
			return err
		}
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

func (a *Api) createGetApi() error {
	// add api file
	handlerFile, err := os.Create(fmt.Sprintf("%s/cmd/%sHandler.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer handlerFile.Close()
	handlerTpl := tpl.GetApiHandlerTemplate()
	apiTemplate := template.Must(template.New("api").Parse(string(handlerTpl)))
	err = apiTemplate.Execute(handlerFile, a)
	if err != nil {
		return err
	}

	// add handler test file
	handlerTestFile, err := os.Create(fmt.Sprintf("%s/cmd/%sHandler_test.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer handlerTestFile.Close()
	handlerTestTpl := tpl.GetApiHandlerTestTemplate()
	handlerTestTemplate := template.Must(template.New("handlerTest").Parse(string(handlerTestTpl)))
	err = handlerTestTemplate.Execute(handlerTestFile, a)
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
	ctrlFile, err := os.Create(fmt.Sprintf("%s/internal/%sCtrl.go", a.ProjectDirectory, a.Name))
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

	// add the test file
	ctrlTestFile, err := os.Create(fmt.Sprintf("%s/internal/%s_test.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer ctrlFile.Close()
	ctrlTestsTpl := tpl.GetApiCtrlTestsTemplate()
	ctrlTestsTemplate := template.Must(template.New("ctrlTest").Parse(string(ctrlTestsTpl)))
	err = ctrlTestsTemplate.Execute(ctrlTestFile, a)
	if err != nil {
		return err
	}

	return nil
}

func (a *Api) createPostApi() error {
	// add api file
	apiFile, err := os.Create(fmt.Sprintf("%s/cmd/%sHandler.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer apiFile.Close()
	apiTpl := tpl.PostApiHandlerTemplate()
	apiTemplate := template.Must(template.New("api").Parse(string(apiTpl)))
	err = apiTemplate.Execute(apiFile, a)
	if err != nil {
		return err
	}

	// checking if pkg/api directory exists
	apiMod := fmt.Sprintf("%s/pkg/api", a.ProjectDirectory)
	if _, err := os.Stat(apiMod); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(apiMod, 0754); err != nil {
			return err
		}
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
	ctrlFile, err := os.Create(fmt.Sprintf("%s/internal/%sCtrl.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer ctrlFile.Close()
	ctrlTpl := tpl.PostApiCtrlTemplate()
	ctrlTemplate := template.Must(template.New("ctrl").Parse(string(ctrlTpl)))
	err = ctrlTemplate.Execute(ctrlFile, a)
	if err != nil {
		return err
	}

	// add request/response file
	reqResFile, err := os.Create(fmt.Sprintf("%s/pkg/api/%sReqRes.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer reqResFile.Close()
	// request / response structs
	reqResTpl := tpl.RequestResponseTemplate()
	reqResTemplate := template.Must(template.New("reqRes").Parse(string(reqResTpl)))
	err = reqResTemplate.Execute(reqResFile, a)
	if err != nil {
		return err
	}

	// add the test file
	ctrlTestFile, err := os.Create(fmt.Sprintf("%s/internal/%s_test.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer ctrlFile.Close()

	ctrlTestsTpl := tpl.PostApiCtrlTestsTemplate()
	ctrlTestsTemplate := template.Must(template.New("ctrlTest").Parse(string(ctrlTestsTpl)))
	err = ctrlTestsTemplate.Execute(ctrlTestFile, a)
	if err != nil {
		return err
	}
	return nil
}

func (a *Api) createPutApi() error {
	// add api file
	apiFile, err := os.Create(fmt.Sprintf("%s/cmd/%sHandler.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer apiFile.Close()
	apiTpl := tpl.PutApiHandlerTemplate()
	apiTemplate := template.Must(template.New("api").Parse(string(apiTpl)))
	err = apiTemplate.Execute(apiFile, a)
	if err != nil {
		return err
	}

	// checking if pkg/api directory exists
	apiMod := fmt.Sprintf("%s/pkg/api", a.ProjectDirectory)
	if _, err := os.Stat(apiMod); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(apiMod, 0754); err != nil {
			return err
		}
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
	ctrlFile, err := os.Create(fmt.Sprintf("%s/internal/%sCtrl.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer ctrlFile.Close()
	ctrlTpl := tpl.PutApiCtrlTemplate()
	ctrlTemplate := template.Must(template.New("ctrl").Parse(string(ctrlTpl)))
	err = ctrlTemplate.Execute(ctrlFile, a)
	if err != nil {
		return err
	}

	// add request/response file
	reqResFile, err := os.Create(fmt.Sprintf("%s/pkg/api/%sReqRes.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer reqResFile.Close()
	// request / response structs
	reqResTpl := tpl.RequestResponseTemplate()
	reqResTemplate := template.Must(template.New("reqRes").Parse(string(reqResTpl)))
	err = reqResTemplate.Execute(reqResFile, a)
	if err != nil {
		return err
	}

	// add the test file
	ctrlTestFile, err := os.Create(fmt.Sprintf("%s/internal/%s_test.go", a.ProjectDirectory, a.Name))
	if err != nil {
		return err
	}
	defer ctrlFile.Close()

	ctrlTestsTpl := tpl.PutApiCtrlTestsTemplate()
	ctrlTestsTemplate := template.Must(template.New("ctrlTest").Parse(string(ctrlTestsTpl)))
	err = ctrlTestsTemplate.Execute(ctrlTestFile, a)
	if err != nil {
		return err
	}
	return nil
}
