package request

import (
	"qiyuesuo/sdk/http"
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
