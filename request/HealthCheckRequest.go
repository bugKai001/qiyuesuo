package request

import (
	"github.com/bugKai001/qiyuesuo/http"
)

type HealthCheckRequest struct {
}

func (obj HealthCheckRequest) GetUrl() string {
	return "/v2/checkhealth"
}

func (obj HealthCheckRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	return parameter
}
