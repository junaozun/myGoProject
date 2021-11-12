package errno

import (
	"encoding/json"
)

type RET interface {
	WithData(data interface{}) RET // WithData 设置成功时返回的数据
	SetMsg(msg string) RET         // 修改返回的默认message
	RetCode() int                  // 返回的错误码
	ToString() string              // ToString 返回 JSON 格式的错误详情
}

type err struct {
	Code int         `json:"code"` // 业务code
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 成功时返回的数据
}

func NewError(code int, msg string) RET {
	return &err{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (e *err) RetCode() int {
	return e.Code
}

func (e err) WithData(data interface{}) RET {
	e.Data = data
	return &e
}

func (e err) SetMsg(msg string) RET {
	e.Msg = msg
	return &e
}

// ToString 返回 JSON 格式的错误详情
func (e *err) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
