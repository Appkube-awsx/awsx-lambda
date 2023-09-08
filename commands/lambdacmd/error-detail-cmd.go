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

// getConfigDataCmd represents the getConfigData command
var GetDetailOfErrorCmd = &cobra.Command{
	Use:   "errorDetail",
	Short: "lambda error details",
	Long:  `get error details of lambda function`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		function, _ := cmd.Flags().GetString("function")

		if authFlag {
			controllers.LambdaDetailsErrorController(function, *clientAuth)

		}

	},
}

func init() {

	GetDetailOfErrorCmd.Flags().StringP("function", "f", "", "lambda function name")

	if err := GetConfigDataCmd.MarkFlagRequired("function"); err != nil {
		fmt.Println("--function is required", err)
	}

}
