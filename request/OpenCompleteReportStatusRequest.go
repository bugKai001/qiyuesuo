package request

import (
	"github.com/bugKai001/qiyuesuo/http"
)

type OpenCompleteReportStatusRequest struct {
	ProveId string `json:"proveId,omitempty"`
}

func (obj OpenCompleteReportStatusRequest) GetUrl() string {
	return "/chain/evidence/complete/report/status"
}

func (obj OpenCompleteReportStatusRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("proveId", obj.ProveId)
	return parameter
}
