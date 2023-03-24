/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetNumberOfErrorCmd = &cobra.Command{
	Use:   "errorCount",
	Short: "to total number of errors",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticater.ChildCommandAuth(cmd)

		if authFlag {
			lambdaClient := client.GetClient()
			cloudClient := client.GetCloudWatchClient()

			function, _ := cmd.Flags().GetString("function")

			if function != "" {
				controllers.GetFunctionErrCount(cloudClient, function)
			} else {
				controllers.GetAllFunctionsErrCount(cloudClient, lambdaClient)
			}
		}

	},
}

func init() {
	GetNumberOfErrorCmd.Flags().StringP("function", "f", "", "lambda function name")
}
