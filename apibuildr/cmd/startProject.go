package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/focks/apibuildr/apibuildr/cmd/tpl"
	"github.com/spf13/cobra"

	"os"
)

type ApibuildrProject struct {
	GoPackage        string
	WorkingDirectory string
	ProjectName      string
}

func (project *ApibuildrProject) Create() error {
	dir := filepath.Join(project.WorkingDirectory, project.ProjectName)
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		// create directory
		CheckError(fmt.Sprintf("directory %s exists", project.ProjectName))
	}
	if err := os.Mkdir(dir, 0754); err != nil {
		return err
	}

	files := []*ApiFile{
		{
			path:     fmt.Sprintf("%s/go.mod", dir),
			template: tpl.GoModTemplate,
		},
	}

	if err := makeFiles(&files, project); err != nil {
		reverseFiles(&files)
		return err
	}

	return nil

}

var startProjectCmd = &cobra.Command{
	Use:   "startProject",
	Short: "creating a new apibuildr project",
	Long:  `create a new project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			CheckError("use apibuildr startproject <project-name> --package packagename")
		}

		wd, err := os.Getwd()
		if err != nil {
			CheckError(err)
		}
		projectName := args[0]

		goPackage, err := cmd.Flags().GetString("package")
		if err != nil {
			CheckError(err)
		}
		if goPackage == "" {
			goPackage = projectName
		}

		project := ApibuildrProject{
			GoPackage:        goPackage,
			WorkingDirectory: wd,
			ProjectName:      projectName,
		}

		if err := project.Create(); err != nil {
			CheckError(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(startProjectCmd)

	startProjectCmd.Flags().StringP("package", "p", "", "package name")
}
