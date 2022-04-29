/**
 * @Author: zze
 * @Date: 2022/4/29 14:59
 * @Desc:
 */
package zvalidator

import (
	"strings"
)

// Required 非空验证
func Required() Validator {
	return &required{
		errorMsg: "不能为空",
	}
}

type required struct {
	errorMsg string
}

func (r *required) validate(value interface{}) bool {
	if value == nil {
		return false
	}
	v, isString := value.(string)
	if isString {
		return len(strings.TrimSpace(v)) > 0
	}
	return true
}

func (r *required) ErrorMsg(msg string) Validator {
	r.errorMsg = msg
	return r
}

func (r *required) getErrorMsg() string {
	return r.errorMsg
}
