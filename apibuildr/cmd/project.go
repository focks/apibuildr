package cmd

import (
	"fmt"
	"os"

	"github.com/focks/apibuildr/apibuildr/cmd/tpl"

	"text/template"
)

type Project struct {
	Name         string
	AbsolutePath string
	Wd           string
	PackageName  string
}

type InitFileVars struct {
	Package string
	Path    string
}

func (p *Project) Create() error {
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	}

	// create main.go
	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer mainFile.Close()

	mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
	err = mainTemplate.Execute(mainFile, p)
	if err != nil {
		return err
	}

	// create cmd/server.go
	if _, err = os.Stat(fmt.Sprintf("%s/cmd", p.AbsolutePath)); os.IsNotExist(err) {
		CheckError(os.Mkdir(fmt.Sprintf("%s/cmd", p.AbsolutePath), 0751))
	}

	pkg := fmt.Sprintf("%s/pkg", p.AbsolutePath)
	if _, err := os.Stat(pkg); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(pkg, 0754); err != nil {
			return err
		}
	}
	rootFile, err := os.Create(fmt.Sprintf("%s/cmd/server.go", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer rootFile.Close()

	rootTemplate := template.Must(template.New("root").Parse(string(tpl.ServerTemplate())))
	err = rootTemplate.Execute(rootFile, p)
	if err != nil {
		return err
	}

	// test utils file
	testUtilsFile, err := os.Create(fmt.Sprintf("%s/cmd/testUtils.go", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer testUtilsFile.Close()
	testUtilsTemplate := template.Must(template.New("testutils").Parse(string(tpl.GetTestUtilTemplate())))
	err = testUtilsTemplate.Execute(testUtilsFile, p)
	if err != nil {
		return err
	}

	return createInitFile(InitFileVars{
		Path:    fmt.Sprintf("%s/cmd", p.AbsolutePath),
		Package: "cmd",
	})

}

func createInitFile(vars InitFileVars) error {
	// create init.go
	initFile, err := os.Create(fmt.Sprintf("%s/init.go", vars.Path))
	if err != nil {
		return err
	}
	defer initFile.Close()

	initTemplate := template.Must(template.New("init").Parse(string(tpl.InitFileTemplate())))
	err = initTemplate.Execute(initFile, vars)
	if err != nil {
		return err
	}

	return nil
}
