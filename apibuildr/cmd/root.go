/*
Copyright © 2022 focks cskkman@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apibuildr",
	Short: "apibuildr generates boilerplate code for Rest REST API handlers",
	Long: `apibuildr is a commandline tool for creating rest apis in golang language.

apibuildr makes it easy to generate boilerplate code while adding rest apis.
	
It is not a framework, it uses gorilla mux server internally.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
