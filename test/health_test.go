package test

import (
	"fmt"
	"qiyuesuo/sdk/request"
	"testing"
)

func TestCheckHealth(t *testing.T) {
	req := request.HealthCheckRequest{}

	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestSaasCheckHealth(t *testing.T) {
	req := request.SaasHealthCheckRequest{}

	response, err := client.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
