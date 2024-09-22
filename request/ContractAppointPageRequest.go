package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type ContractAppointPageRequest struct {
	// 合同ID和业务ID不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID和业务ID不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID获取签署页面，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 指定签署位置页面：APPOINT，参数填写页面：FILL_PARAMETERS，基础信息编辑页面：DRAFT，默认为指定签署位置页面
	PageType     string `json:"pageType,omitempty"`
	CallbackPage string `json:"callbackPage,omitempty"`
	// 默认允许发起
	AllowedSend *bool `json:"allowedSend,omitempty"`
	// 默认false允许显示返回按钮
	HideReturnButton *bool `json:"hideReturnButton,omitempty"`
	// cn（中文）、en（英文）；默认以发起人语言为准，合同中未传发起人时，展示中文
	Lang      string           `json:"lang,omitempty"`
	PageStyle *model.PageStyle `json:"pageStyle,omitempty"`
	// 默认为true展示按钮
	CanSaveOrClose        *bool `json:"canSaveOrClose,omitempty"`
	SpecifyLocationButton *bool `json:"specifyLocationButton,omitempty"`
	FillParamButton       *bool `json:"fillParamButton,omitempty"`
}

func (obj ContractAppointPageRequest) GetUrl() string {
	return "/v2/contract/appointurl"
}

func (obj ContractAppointPageRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}

/*
原类名写错，不过为了兼容旧版本，依旧保留；
后续使用上面新加的类，此类不再更新。
*/
type ContractAppointurlRequest struct {
	// 合同ID和业务ID不能同时为空
	ContractId string `json:"contractId,omitempty"`
	// 合同ID和业务ID不能同时为空
	BizId string `json:"bizId,omitempty"`
	// 若使用业务ID获取签署页面，且合同是以子公司身份创建的，则需要传递该值，用于确定合同主体
	TenantName string `json:"tenantName,omitempty"`
	// 指定签署位置页面：APPOINT，参数填写页面：FILL_PARAMETERS，基础信息编辑页面：DRAFT，默认为指定签署位置页面
	PageType     string `json:"pageType,omitempty"`
	CallbackPage string `json:"callbackPage,omitempty"`
	// 默认允许发起
	AllowedSend *bool `json:"allowedSend,omitempty"`
	// 默认false允许显示返回按钮
	HideReturnButton *bool `json:"hideReturnButton,omitempty"`
	// cn（中文）、en（英文）；默认以发起人语言为准，合同中未传发起人时，展示中文
	Lang      string           `json:"lang,omitempty"`
	PageStyle *model.PageStyle `json:"pageStyle,omitempty"`
	// 默认为true展示按钮
	CanSaveOrClose        *bool `json:"canSaveOrClose,omitempty"`
	SpecifyLocationButton *bool `json:"specifyLocationButton,omitempty"`
	FillParamButton       *bool `json:"fillParamButton,omitempty"`
}

func (obj ContractAppointurlRequest) GetUrl() string {
	return "/v2/contract/appointurl"
}

func (obj ContractAppointurlRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
