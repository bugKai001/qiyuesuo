package model

type TextFields struct {
	Id        string  `json:"id,omitempty"`        // 文本域ID
	Width     int     `json:"width,omitempty"`     // 文本域宽度
	Height    int     `json:"height,omitempty"`    // 文本域高度
	Keyword   string  `json:"keyword,omitempty"`   // 关键字
	LlyOffset int     `json:"llyOffset,omitempty"` // 文本域左下角y轴坐标相对于关键字左下角y轴坐标的偏移量
	LlxOffset int     `json:"llxOffset,omitempty"` // 文本域左下角x轴坐标相对于关键字左下角x轴坐标的偏移量
	Lly       int     `json:"lly,omitempty"`       // 文本域左下角y轴坐标相对于对于页面的偏移比例
	Llx       int     `json:"llx,omitempty"`       // 文本域左下角x轴坐标相对于页面的偏移比例
	FontSize  int     `json:"fontSize,omitempty"`  // 文本域字体大小
	PageIndex float32 `json:"pageIndex,omitempty"` // 文本域所在页码
}
