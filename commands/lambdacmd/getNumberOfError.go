/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetNumberOfErrorCmd = &cobra.Command{
	Use:   "errorCount",
	Short: "to total number of errors",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth := authenticater.ChildCommandAuth(cmd)
		function, _ := cmd.Flags().GetString("function")

		if authFlag {
			controllers.LambdaGetNumberOfErrorController(function, *clientAuth)

		}

	},
}

func init() {
	GetNumberOfErrorCmd.Flags().StringP("function", "f", "", "lambda function name")
}
