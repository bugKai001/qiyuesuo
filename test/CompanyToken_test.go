package test

import (
	"fmt"
	"github.com/bugKai001/qiyuesuo/request"
	"testing"
)

func TestCompanyTokenDetail(t *testing.T) {
	req := request.CompanyTokenDetailRequest{}
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}

func TestCompanyTokenUpdate(t *testing.T) {
	req := request.CompanyTokenUpdateCallbackRequest{}
	req.CallbackUrl = "https://test.qiyuesuo.me"
	encrypt := true
	req.Encrypt = &encrypt
	response, err := sdkClient.Service(req)
	if err != nil {
		fmt.Println("request failed,", err.Error())
		return
	}
	fmt.Println(response)
}
