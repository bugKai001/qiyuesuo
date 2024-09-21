package request

import (
	"qiyuesuo/sdk/http"
)

type SaasHealthCheckRequest struct {
}

func (obj SaasHealthCheckRequest) GetUrl() string {
	return "/saas/v2/checkhealth"
}

func (obj SaasHealthCheckRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	return parameter
}
