package http

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/commons"
	"io"
	"strings"
)

const (
	SDK_VERSION = "GO-3.0.5"
)

type SdkClient struct {
	ServerUrl    string
	AccessToken  string
	AccessSecret string
	SdkVersion   string
	Language     string
}

func NewSdkClient(serverUrl string, accessToken string, accessSecret string) *SdkClient {
	sdkClient := SdkClient{
		ServerUrl:    serverUrl,
		AccessToken:  accessToken,
		AccessSecret: accessSecret,
		SdkVersion:   SDK_VERSION,
	}
	return &sdkClient
}

func NewSdkClient2(serverUrl string, accessToken string, accessSecret string, sdkVersion string) *SdkClient {
	sdkClient := SdkClient{
		ServerUrl:    serverUrl,
		AccessToken:  accessToken,
		AccessSecret: accessSecret,
		SdkVersion:   sdkVersion,
	}
	return &sdkClient
}

func (sdk *SdkClient) Service(request SdkRequest) (res map[string]interface{}, err error) {
	bs, err := sdk.sdkRequest(request)
	res = make(map[string]interface{})
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &res)
	return
}

func (sdk *SdkClient) Download(request SdkRequest, w io.Writer) (err error) {
	getUrl := request.GetUrl()
	if strings.HasPrefix(getUrl, "/v2/") {
		getUrl = strings.Replace(getUrl, "/v2", "", 1)
	}
	url := sdk.ServerUrl + getUrl
	parameter := request.GetHttpParameter()
	header := sdk.buildHttpHeader(parameter)
	if parameter.IsJson() {
		err = DoDownloadWithJson(url, parameter, header, w)
	} else {
		err = DoDownload(url, parameter, header, w)
	}
	return
}

func (sdk *SdkClient) sdkRequest(request SdkRequest) (bs []byte, err error) {
	getUrl := request.GetUrl()
	if strings.HasPrefix(getUrl, "/v2/") {
		getUrl = strings.Replace(getUrl, "/v2", "", 1)
	}
	url := sdk.ServerUrl + getUrl
	parameter := request.GetHttpParameter()
	header := sdk.buildHttpHeader(parameter)
	if parameter.IsJson() {
		return DoServiceWithJson(url, parameter, header)
	}
	return DoService(url, parameter, header)
}

func (sdk *SdkClient) buildHttpHeader(httpParameter *HttpParameter) *HttpHeader {
	timestamp := commons.GetTimeStamp()
	nonce := commons.GetUUID()
	var signature = commons.GetMD5(sdk.AccessToken + sdk.AccessSecret + timestamp + nonce)
	header := HttpHeader{
		AccessToken: sdk.AccessToken,
		Timestamp:   timestamp,
		Nonce:       nonce,
		Signature:   signature,
		SdkVersion:  sdk.SdkVersion,
		Language:    sdk.Language,
	}
	return &header
}
