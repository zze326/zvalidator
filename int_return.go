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

type intReturn struct {
	valuePointer *int
	errorMsg     string
}

// IntReturn 整型数字校验，验证成功后结果将传入 Int 类型指针
func IntReturn(valuePointer *int) Validator {
	return &intReturn{
		errorMsg: "必须是整型数字",
	}
}

func (i *intReturn) validate(value interface{}) bool {
	v, err := strconv.Atoi(fmt.Sprintf("%v", value))
	if err != nil {
		return false
	}
	i.valuePointer = &v
	return true
}

func (i *intReturn) ErrorMsg(msg string) Validator {
	i.errorMsg = msg
	return i
}

func (i *intReturn) getErrorMsg() string {
	return i.errorMsg
}
