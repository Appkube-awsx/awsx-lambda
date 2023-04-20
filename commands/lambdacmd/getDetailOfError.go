/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/client"
	"github.com/Appkube-awsx/awsx-lambda/services"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetDetailOfErrorCmd = &cobra.Command{
	Use:   "errorDetail",
	Short: "to get details of error in lambda",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticater.ChildCommandAuth(cmd)

		if authFlag {
			cloudClient := client.GetCloudWatchClient()
			function, _ := cmd.Flags().GetString("function")
			services.GetFunctionsErrDetail(cloudClient, function)
		}

	},
}

func init() {

	GetDetailOfErrorCmd.Flags().StringP("function", "f", "", "lambda function name")

	if err := GetConfigDataCmd.MarkFlagRequired("function"); err != nil {
		fmt.Println("--function is required", err)
	}

}
