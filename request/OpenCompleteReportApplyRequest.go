package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type OpenCompleteReportApplyRequest struct {
	ProveSource   string                 `json:"proveSource,omitempty"`
	ChainId       string                 `json:"chainId,omitempty"`
	ProveDataId   string                 `json:"proveDataId,omitempty"`
	DocumentIds   string                 `json:"documentIds,omitempty"`
	OperatorInfo  *model.UserInfoRequest `json:"operatorInfo,omitempty"`
	ProveMaterial []*http.FileItem       `json:"proveMaterial,omitempty"`
}

func (obj OpenCompleteReportApplyRequest) GetUrl() string {
	return "/chain/evidence/complete/report/apply"
}

func (obj OpenCompleteReportApplyRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("proveSource", obj.ProveSource)
	parameter.AddParam("chainId", obj.ChainId)
	parameter.AddParam("proveDataId", obj.ProveDataId)
	parameter.AddParam("documentIds", obj.DocumentIds)
	if obj.OperatorInfo != nil {
		jsonBytes, _ := json.Marshal(obj.OperatorInfo)
		parameter.AddParam("operatorInfo", string(jsonBytes))
	}
	parameter.AddListFiles("proveMaterial", obj.ProveMaterial)
	return parameter
}
