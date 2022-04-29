/**
 * @Author: zze
 * @Date: 2022/4/29 15:51
 * @Desc:
 */
package zvalidator

import (
	"fmt"
	"strconv"
)

type lessThanEqual struct {
	errorMsg string
	target   float64
}

// LessThanEqual 验证是否小于
func LessThanEqual(target float64) Validator {
	return &lessThanEqual{
		errorMsg: "不能小于等于 " + strconv.FormatFloat(target, 'f', -1, 64),
		target:   target,
	}
}

func (l *lessThanEqual) validate(value interface{}) bool {
	result, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	if err != nil {
		l.errorMsg = fmt.Sprintf("%v 不是一个数字", value)
		return false
	}
	return result <= l.target
}

func (l *lessThanEqual) ErrorMsg(msg string) Validator {
	l.errorMsg = msg
	return l
}

func (l *lessThanEqual) getErrorMsg() string {
	return l.errorMsg
}
