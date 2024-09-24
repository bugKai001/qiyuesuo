package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

// DocAddByTemplateRequest 用模板创建合同文档
type DocAddByTemplateRequest struct {
	Title      string              `json:"title,omitempty"`      // 合同文档名称,默认使用模板名称
	TemplateId string              `json:"templateId,omitempty"` // 模板ID
	Params     string              `json:"params,omitempty"`     // 模板参数，json格式字符串 如：{"param1":"value1","param2":"value2"}
	WaterMarks []*model.Watermark  `json:"waterMarks,omitempty"` // 水印设置
	Form       []*model.TextFields `json:"form,omitempty"`       // 模板表单设置
}

func (obj DocAddByTemplateRequest) GetUrl() string {
	return "/document/createbytemplate"
}

func (obj DocAddByTemplateRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
