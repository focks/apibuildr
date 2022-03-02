package cmd

import (
	"fmt"

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
		apiPath, err := cmd.Flags().GetString("path")
		if err != nil {
			CheckError(err)
		}
		if apiPath == "" {
			CheckError("path cannot be empty , enter a path with -p flag")
		}

		modName := getModImportPath()
		apiPath = strings.Trim(apiPath, "/")
		splits := strings.Split(apiPath, "/")
		pathEnd := splits[len(splits)-1]
		apiPath = strings.Trim(apiPath, fmt.Sprintf("/%s", pathEnd))

		api := Api{
			Name:             strings.Title(apiName),
			Method:           apiMethod,
			Path:             apiPath,
			PathEnd:          pathEnd,
			PackageName:      modName,
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
