/*
	Copyright © 2022 focks cskkman@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize the application",
	Long: `
Initialize (apiBuildr init) will create a new application/executable, with a license
and the appropriate structure for a gorilla mux based REST API application.
apibuildr init must be run inside of a go module (please run "go mod init <MODNAME>" first)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := initializeProject(args)
		CheckError(err)
		fmt.Println("Your application is ready.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initializeProject(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = fmt.Sprintf("%s/%s", wd, args[0])
		}
	}

	modName := getModImportPath()

	project := &Project{
		AbsolutePath: wd,
		PackageName:  modName,
		Name:         path.Base(modName),
	}

	if err := project.Create(); err != nil {
		return err
	}

	CheckError(goGet("github.com/focks/apibuildr"))
	return nil
}
