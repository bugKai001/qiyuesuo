package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type OpenChainEvidenceProveRequest struct {
	ChainId      string                 `json:"chainId,omitempty"`
	ProveDataId  string                 `json:"proveDataId,omitempty"`
	ProveSource  string                 `json:"proveSource,omitempty"`
	OperatorInfo *model.UserInfoRequest `json:"operatorInfo,omitempty"`
}

func (obj OpenChainEvidenceProveRequest) GetUrl() string {
	return "/chain/evidence/electronic/prove/apply"
}

func (obj OpenChainEvidenceProveRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("chainId", obj.ChainId)
	parameter.AddParam("proveDataId", obj.ProveDataId)
	parameter.AddParam("proveSource", obj.ProveSource)
	if obj.OperatorInfo != nil {
		jsonBytes, _ := json.Marshal(obj.OperatorInfo)
		parameter.AddParam("operatorInfo", string(jsonBytes))
	}
	return parameter
}
