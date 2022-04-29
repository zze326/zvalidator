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

type greaterThanEqual struct {
	errorMsg string
	target   float64
}

// GreaterThanEqual 验证是否大于等于
func GreaterThanEqual(target float64) Validator {
	return &greaterThanEqual{
		errorMsg: "必须大于等于 " + strconv.FormatFloat(target, 'f', -1, 64),
		target:   target,
	}
}

func (g *greaterThanEqual) validate(value interface{}) bool {
	result, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	if err != nil {
		g.errorMsg = fmt.Sprintf("%v 不是一个数字", value)
		return false
	}
	return result >= g.target
}

func (g *greaterThanEqual) ErrorMsg(msg string) Validator {
	g.errorMsg = msg
	return g
}

func (g *greaterThanEqual) getErrorMsg() string {
	return g.errorMsg
}
