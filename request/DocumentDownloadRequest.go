package request

import (
	"github.com/bugKai001/qiyuesuo/http"
)

type DocumentDownloadRequest struct {
	// 合同文档ID
	DocumentId string `json:"documentId,omitempty"`
}

func (obj DocumentDownloadRequest) GetUrl() string {
	return "/v2/document/download"
}

func (obj DocumentDownloadRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("documentId", obj.DocumentId)
	return parameter
}
