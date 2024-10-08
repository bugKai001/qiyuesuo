package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type CompanyAuthMiniappTicketRequest struct {
	CompanyName string      `json:"companyName,omitempty"`
	Applicant   *model.User `json:"applicant,omitempty"`
	CallbackUrl string      `json:"callbackUrl,omitempty"`
}

func (obj CompanyAuthMiniappTicketRequest) GetUrl() string {
	return "/companyauth/miniappexchange"
}

func (obj CompanyAuthMiniappTicketRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
