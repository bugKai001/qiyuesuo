package request

import (
	"github.com/bugKai001/qiyuesuo/http"
)

type PlatformDetailRequest struct {
}

func (obj PlatformDetailRequest) GetUrl() string {
	return "/company/platforminfo"
}

func (obj PlatformDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	return parameter
}
