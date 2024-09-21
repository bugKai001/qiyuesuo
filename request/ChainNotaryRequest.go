package request

import (
	"encoding/json"
	"qiyuesuo/sdk/http"
	"qiyuesuo/sdk/model"
)

type ChainNotaryRequest struct {
	NotaryName     string                 `json:"notaryName,omitempty"`
	NotaryDataID   string                 `json:"notaryDataID,omitempty"`
	NotaryType     string                 `json:"notaryType,omitempty"`
	FileType       string                 `json:"fileType,omitempty"`
	NotaryFile     *http.FileItem         `json:"notaryFile,omitempty"`
	NotaryDataHash string                 `json:"notaryDataHash,omitempty"`
	HashAlgorithm  string                 `json:"hashAlgorithm,omitempty"`
	OperatorInfo   *model.UserInfoRequest `json:"operatorInfo,omitempty"`
}

func (obj ChainNotaryRequest) GetUrl() string {
	return "/v2/chain/notary"
}

func (obj ChainNotaryRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	parameter.AddParam("notaryName", obj.NotaryName)
	parameter.AddParam("notaryDataID", obj.NotaryDataID)
	parameter.AddParam("notaryType", obj.NotaryType)
	parameter.AddParam("fileType", obj.FileType)
	parameter.AddFiles("notaryFile", obj.NotaryFile)
	parameter.AddParam("notaryDataHash", obj.NotaryDataHash)
	parameter.AddParam("hashAlgorithm", obj.HashAlgorithm)
	if obj.OperatorInfo != nil {
		jsonBytes, _ := json.Marshal(obj.OperatorInfo)
		parameter.AddParam("operatorInfo", string(jsonBytes))
	}
	return parameter
}
