package code

import (
	"fmt"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 公共返回信息
var Success = &Result{Code: 0, Msg: "操作成功", Data: nil}

var Fail = &Result{Code: 1, Msg: "操作失败", Data: nil}

type Err struct {
	Msg   string
	Cause error
}

func (e *Err) Error() string {
	if e.Cause == nil {
		return e.Msg
	}

	return fmt.Sprintf("%s: %s", e.Msg, e.Cause.Error())
}

func New(msg string) error {
	return &Err{
		Msg:   msg,
		Cause: nil,
	}
}

func WithMessage(err error, msg string) error {
	return &Err{
		Msg:   msg,
		Cause: err,
	}
}
