package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"strings"

	"os"
)

// addApiCmd represents the addApi command
var addApiCmd = &cobra.Command{
	Use:   "addApi",
	Short: "adding a new api",
	Long:  `add new api`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			CheckError("use apibuildr addApi <api-name> --method GET --path /v1/path")
		}

		wd, err := os.Getwd()
		if err != nil {
			CheckError(err)
		}
		apiName := args[0]

		apiMethod, err := cmd.Flags().GetString("method")
		if err != nil {
			CheckError(err)
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			CheckError(err)
		}

		modName := getModImportPath()

		api := Api{
			Name:             strings.Title(apiName),
			Method:           apiMethod,
			Path:             path,
			Uri:              makeUriPath(path),
			PackageName:      modName,
			ProjectDirectory: wd,
		}

		log.Println(api.String())

		if err := api.Create(); err != nil {
			CheckError(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(addApiCmd)

	addApiCmd.Flags().StringP("path", "p", "", "api endpoint path")
	addApiCmd.Flags().StringP("method", "m", "GET", "http method")
}
