/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// GetConfigWithTagsCmd represents the getConfigWithTags command
var GetConfigWithTagsCmd = &cobra.Command{
	Use:   "getConfigWithTags",
	Short: "lambda configuration with tags",
	Long:  `lambda configuration with tags`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		function, _ := cmd.Flags().GetString("function")

		if authFlag {
			controllers.LambdaFunctionWithTagsController(function, *clientAuth)
		}
	},
}

func init() {
	GetConfigWithTagsCmd.Flags().StringP("function", "f", "", "lambda function name")

	if err := GetConfigWithTagsCmd.MarkFlagRequired("function"); err != nil {
		fmt.Println("--function or -f is required", err)
	}
}
