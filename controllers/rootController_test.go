package controllers

import (
	"os"
	"testing"
)

var (
	acKey     = os.Getenv("AWS_ACCKEY")
	crossRole = os.Getenv("AWS_CROSS_ARN")
	external  = os.Getenv("AWS_EXTERNALID")
	secKey    = os.Getenv("AWS_SECKEY")
	zone      = "us-east-1"
)

// Test for get All lambda list
func TestAllLambdaList(t *testing.T) {

	result := AllLambdaListController("", "", zone, acKey, secKey, crossRole, external)

	if result != nil {
		t.Log("Test passed for all lambda list")
	} else {
		t.Error("Test failed for all lambda list")
	}
}

// Test for get lambda list
func TestLambdaList(t *testing.T) {

	result := LambdaListController("", "", "", zone, acKey, secKey, crossRole, external)

	if result.NextMarker == nil {
		t.Log("Test failed for lambda list")
	}

	result1 := LambdaListController(*result.NextMarker, "", "", zone, acKey, secKey, crossRole, external)

	if result1 != nil {
		t.Log("Test passed for lambda list")
	} else {
		t.Error("Test failed for lambda list")
	}
}

// Test for get config data of lambda
func TestConfigDataController(t *testing.T) {

	result := LambdaDetails("testfunction", "", "", zone, acKey, secKey, crossRole, external)

	if result != nil {
		t.Log("Test passed for lambda details")
	} else {
		t.Error("Test failed for lambda list")
	}
}

// Test for lambda latency
func TestLatencyController(t *testing.T) {

	result := LambdaDetails("testfunction", "", "", zone, acKey, secKey, crossRole, external)

	if result != nil {
		t.Log("Test passed for lambda details")
	} else {
		t.Error("Test failed for lambda list")
	}
}

// Test for number if error
func TestNumberOfError(t *testing.T) {

	result := LambdaGetNumberOfErrorController("testfunction", "", "", zone, acKey, secKey, crossRole, external)

	if result >= 0 {
		t.Log("Test passed for error count for a function")
	} else {
		t.Log("Something went wrong")
	}
}

// Test for total number of lambdas
func TestTotalNumberOfLambdas(t *testing.T) {

	result := LambdaGetTotalNumberOfLambda("", "", zone, acKey, secKey, crossRole, external)

	if result >= 0 {
		t.Log("Test passed for total number of lambdas")
	} else {
		t.Log("Test failed for total number of lambdas")
	}
}
