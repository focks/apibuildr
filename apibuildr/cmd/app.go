package cmd

import (
	"fmt"
	"os"

	"github.com/focks/apibuildr/apibuildr/cmd/tpl"
)

type App struct {
	Name         string
	AppDirectory string
	Wd           string
	PackageName  string
}

func (app *App) Create() error {
	if _, err := os.Stat(app.AppDirectory); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(app.AppDirectory, 0754); err != nil {
			return err
		}
	}

	if _, err := os.Stat(fmt.Sprintf("%s/cmd", app.AppDirectory)); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(fmt.Sprintf("%s/cmd", app.AppDirectory), 0754); err != nil {
			return err
		}
	}

	if _, err := os.Stat(fmt.Sprintf("%s/pkg", app.AppDirectory)); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(fmt.Sprintf("%s/pkg", app.AppDirectory), 0754); err != nil {
			return err
		}
	}

	files := []*ApiFile{
		{
			path:     fmt.Sprintf("%s/main.go", app.AppDirectory),
			template: string(tpl.MainTemplate()),
		},
		{
			path:     fmt.Sprintf("%s/cmd/server.go", app.AppDirectory),
			template: string(tpl.ServerTemplate()),
		},
		{
			path:     fmt.Sprintf("%s/cmd/testUtils.go", app.AppDirectory),
			template: string(tpl.GetTestUtilTemplate()),
		},

		{
			path:     fmt.Sprintf("%s/cmd/init.go", app.AppDirectory),
			template: initFileTemplate,
		},
		{
			path:     fmt.Sprintf("%s/Dockerfile", app.AppDirectory),
			template: tpl.DockerfileTemplate,
		},
		{
			path:     fmt.Sprintf("%s/makefile", app.AppDirectory),
			template: tpl.MakefileTemplate,
		},
	}

	if err := makeFiles(&files, app); err != nil {
		reverseFiles(&files)
		return err
	}

	return nil

}

var initFileTemplate = `package cmd

import "go.uber.org/zap"

var logger *zap.Logger

func setLogger(lg *zap.Logger) {
	logger = lg
}

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	setLogger(l)
}

`
