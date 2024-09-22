package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
)

type ContractMiniappViewTicketRequest struct {
	ContractId string `json:"contractId,omitempty"`
	BizId      string `json:"bizId,omitempty"`
	DocumentId string `json:"documentId,omitempty"`
	TenantName string `json:"tenantName,omitempty"`
	Lang       string `json:"lang,omitempty"`
	ThemeColor string `json:"themeColor,omitempty"`
}

func (obj ContractMiniappViewTicketRequest) GetUrl() string {
	return "/v2/contract/miniappexchange/view"
}

func (obj ContractMiniappViewTicketRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
