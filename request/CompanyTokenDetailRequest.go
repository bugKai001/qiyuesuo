package request

import (
	"qiyuesuo/sdk/http"
)

type CompanyTokenDetailRequest struct {
}

func (obj CompanyTokenDetailRequest) GetUrl() string {
	return "/company/token/get"
}

func (obj CompanyTokenDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	return parameter
}
