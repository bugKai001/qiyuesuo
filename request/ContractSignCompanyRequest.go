package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type ContractSignCompanyRequest struct {
	Param *model.SignParam
}

func (obj ContractSignCompanyRequest) GetUrl() string {
	return "/v2/contract/companysign"
}

func (obj ContractSignCompanyRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj.Param)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
