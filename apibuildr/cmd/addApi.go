package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

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
		_, err = os.Stat(fmt.Sprintf("%s/main.go", wd))
		if err != nil {

			CheckError("the current directory is not apibuildr project.")
		}
		_, err = os.Stat(fmt.Sprintf("%s/cmd/server.go", wd))
		if err != nil {
			CheckError("the current directory is not apibuildr project.")
		}
		apiName := args[0]
		_, err = os.Stat(fmt.Sprintf("%s/cmd/%s.go", wd, apiName))
		if err == nil {
			CheckError("the given api name already exists.")
		}
		apiMethod, err := cmd.Flags().GetString("method")
		if err != nil {
			CheckError(err)
		}
		apiPath, err := cmd.Flags().GetString("path")
		if err != nil {
			CheckError(err)
		}

		api := Api{
			Name:             apiName,
			Method:           apiMethod,
			Path:             apiPath,
			ProjectDirectory: wd,
		}

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

func addApi(apiName, method string) error {
	return nil
}
