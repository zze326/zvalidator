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

type greaterThan struct {
	errorMsg string
	target   float64
}

// GreaterThan 验证是否大于
func GreaterThan(target float64) Validator {
	return &greaterThan{
		errorMsg: "必须大于 " + strconv.FormatFloat(target, 'f', -1, 64),
		target:   target,
	}
}



func (g *greaterThan) validate(value interface{}) bool {
	result, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	if err != nil {
		g.errorMsg = fmt.Sprintf("%v 不是一个数字", value)
		return false
	}
	return result > g.target
}

func (g *greaterThan) ErrorMsg(msg string) Validator {
	g.errorMsg = msg
	return g
}

func (g *greaterThan) getErrorMsg() string {
	return g.errorMsg
}
