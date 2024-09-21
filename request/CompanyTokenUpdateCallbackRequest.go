package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
)

type CompanyTokenUpdateCallbackRequest struct {
	CallbackUrl string `json:"callbackUrl,omitempty"`
	Encrypt     *bool  `json:"encrypt,omitempty"`
}

func (obj CompanyTokenUpdateCallbackRequest) GetUrl() string {
	return "/company/token/update/callback"
}

func (obj CompanyTokenUpdateCallbackRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
