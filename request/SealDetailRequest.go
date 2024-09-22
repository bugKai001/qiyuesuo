package request

import (
	"github.com/bugKai001/qiyuesuo/http"
)

type SealDetailRequest struct {
	// 印章ID
	SealId string `json:"sealId,omitempty"`
}

func (obj SealDetailRequest) GetUrl() string {
	return "/v2/seal/sealdetail"
}

func (obj SealDetailRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("sealId", obj.SealId)
	return parameter
}
