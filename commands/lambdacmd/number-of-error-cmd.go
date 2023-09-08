/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetNumberOfErrorCmd = &cobra.Command{
	Use:   "errorCount",
	Short: "total number of errors",
	Long:  `get total number of errors of lambda function`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		function, _ := cmd.Flags().GetString("function")

		if authFlag {
			controllers.LambdaGetNumberOfErrorController(function, *clientAuth)

		}

	},
}

func init() {
	GetNumberOfErrorCmd.Flags().StringP("function", "f", "", "lambda function name")
}
