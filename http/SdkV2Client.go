package http

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/commons"
	"io"
)

type SdkV2Client struct {
	ServerUrl    string
	AccessToken  string
	AccessSecret string
	SdkVersion   string
	Language     string
}

func NewSdkV2Client(serverUrl string, accessToken string, accessSecret string) *SdkV2Client {
	sdkClient := SdkV2Client{
		ServerUrl:    serverUrl,
		AccessToken:  accessToken,
		AccessSecret: accessSecret,
		SdkVersion:   SDK_VERSION,
	}
	return &sdkClient
}

func (sdk *SdkV2Client) Service(request SdkRequest) (res map[string]interface{}, err error) {
	bs, err := sdk.sdkRequest(request)
	res = make(map[string]interface{})
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &res)
	return
}

func (sdk *SdkV2Client) Download(request SdkRequest, w io.Writer) (err error) {
	url := sdk.ServerUrl + request.GetUrl()
	parameter := request.GetHttpParameter()
	header := sdk.buildHttpHeader(parameter)
	if parameter.IsJson() {
		err = DoDownloadWithJson(url, parameter, header, w)
	} else {
		err = DoDownload(url, parameter, header, w)
	}
	return
}

func (sdk *SdkV2Client) sdkRequest(request SdkRequest) (bs []byte, err error) {
	url := sdk.ServerUrl + request.GetUrl()
	parameter := request.GetHttpParameter()
	header := sdk.buildHttpHeader(parameter)
	if parameter.IsJson() {
		return DoServiceWithJson(url, parameter, header)
	}
	return DoService(url, parameter, header)
}

func (sdk *SdkV2Client) buildHttpHeader(httpParameter *HttpParameter) *HttpHeader {
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
