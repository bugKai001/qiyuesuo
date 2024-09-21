package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type SignatoryOperatorsEditRequest struct {
	ContractId      string        `json:"contractId,omitempty"`
	BizId           string        `json:"bizId,omitempty"`
	ActionId        string        `json:"actionId,omitempty"`
	ActionOperators []*model.User `json:"actionOperators,omitempty"`
	EditType        string        `json:"editType,omitempty"`
	Operator        *model.User   `json:"operator,omitempty"`
}

func (obj SignatoryOperatorsEditRequest) GetUrl() string {
	return "/v2/signatory/operators/edit"
}

func (obj SignatoryOperatorsEditRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
