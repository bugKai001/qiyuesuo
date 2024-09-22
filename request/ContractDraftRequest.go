package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type ContractDraftRequest struct {
	// 接口返回值
	Contract *model.Contract
}

func (obj ContractDraftRequest) GetUrl() string {
	return "/v2/contract/draft"
}

func (obj ContractDraftRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj.Contract)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
