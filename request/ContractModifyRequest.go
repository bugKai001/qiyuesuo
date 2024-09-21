package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ContractModifyRequest struct {
	// 合同ID和业务ID不能同时为空
	Id string `json:"id,omitempty"`
	// 合同ID和业务ID不能同时为空
	BizId   string `json:"bizId,omitempty"`
	Subject string `json:"subject,omitempty"`
	// 若使用业务ID获取签署页面，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	ExpireTime string `json:"expireTime,omitempty"`
	// 格式为yyyy-MM-dd HH:mm:ss， 传递后会重置为当天23:59:59
	EndTime        string                 `json:"endTime,omitempty"`
	Documents      []string               `json:"documents,omitempty"`
	TemplateParams []*model.TemplateParam `json:"templateParams,omitempty"`
}

func (obj ContractModifyRequest) GetUrl() string {
	return "/v2/contract/contractModify"
}

func (obj ContractModifyRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
