/*
	Copyright © 2022 focks cskkman@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var addAppCmd = &cobra.Command{
	Use:   "addApp",
	Short: "start a new microservice",
	Long: `
		New Mircorservice
	`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("add app command", args)
		err := initializeApp(args)
		CheckError(err)
		fmt.Println("Your application is ready.")
	},
}

func init() {
	rootCmd.AddCommand(addAppCmd)
}

func initializeApp(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if len(args) <= 0 {
		CheckError("app name is required")
	}

	if args[0] == "" {
		CheckError("app is not set")
	}

	modName := getModImportPath()

	log.Println("the package name is", filepath.Join(args[0], modName))

	app := &App{
		AppDirectory: filepath.Join(wd, args[0]),
		Wd:           wd,
		PackageName:  filepath.Join(modName, args[0]),
		// Name:         path.Base(modName),
		Name: args[0],
	}

	if err := app.Create(); err != nil {
		return err
	}

	CheckError(goGet("github.com/focks/apibuildr"))
	return nil
}
