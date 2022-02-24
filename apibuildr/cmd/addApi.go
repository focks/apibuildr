package cmd

import (
	"os"

	"github.com/spf13/cobra"
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
		_, err = os.Open("main.go")
		if err != nil {
			CheckError("the current directory doesn't belong to apibuildr environment.")
		}
		_, err = os.Open("cmd/server.go")
		if err != nil {
			CheckError("the current directory doesn't belong to apibuildr environment.")
		}
		apiName := args[0]
		_, err = os.Open("./cmd/" + apiName + ".go")
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
