package request

import (
	"github.com/bugKai001/qiyuesuo/http"
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
