/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package lambdacmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-lambda/authenticater"
	"github.com/Appkube-awsx/awsx-lambda/controllers"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetLatencyCmd = &cobra.Command{
	Use:   "latency",
	Short: "to get latency of lambda function",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticater.ChildCommandAuth(cmd)
		function, _ := cmd.Flags().GetString("function")
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")

		if authFlag {

			controllers.GetLambadaLatencyTimeController(function, startTime, endTime, authenticater.VaultUrl, authenticater.AccountId, authenticater.Region, authenticater.AcKey, authenticater.SecKey, authenticater.CrossAccountRoleArn, authenticater.ExternalId)

			// services.GetLambadaLatencyTime(cloudClient, function, startTime, endTime)

		}

	},
}

func init() {

	GetLatencyCmd.Flags().StringP("function", "f", "", "lambda function name")
	GetLatencyCmd.Flags().StringP("startTime", "s", "", "lambda start Time")
	GetLatencyCmd.Flags().StringP("endTime", "e", "", "lambda endtime")

	if err := GetLatencyCmd.MarkFlagRequired("function"); err != nil {
		fmt.Println("--function is required", err)
	}

}
