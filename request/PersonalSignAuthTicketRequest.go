package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type PersonalSignAuthTicketRequest struct {
	User               *model.User                     `json:"user,omitempty"`
	AuthDeadline       string                          `json:"authDeadline,omitempty"`
	Remark             string                          `json:"remark,omitempty"`
	CallbackUrl        string                          `json:"callbackUrl,omitempty"`
	TenantName         string                          `json:"tenantName,omitempty"`
	AuthMethod         string                          `json:"authMethod,omitempty"`
	AuthInfo           *model.PersonalSignUserAuthInfo `json:"authInfo,omitempty"`
	PageStyle          *model.PageStyle                `json:"pageStyle,omitempty"`
	AllowPersonalAuto  *bool                           `json:"allowPersonalAuto,omitempty"`
	AuthTimeModifiable *bool                           `json:"authTimeModifiable,omitempty"`
}

func (obj PersonalSignAuthTicketRequest) GetUrl() string {
	return "/v2/personalsign/authurl/miniappexchange"
}

func (obj PersonalSignAuthTicketRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
