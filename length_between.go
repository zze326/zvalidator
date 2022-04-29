/**
 * @Author: zze
 * @Date: 2022/4/29 16:47
 * @Desc:
 */
package zvalidator

import (
	"fmt"
)

type lengthBetween struct {
	errorMsg string
	start    int
	end      int
}

// LengthBetween 长度范围
func LengthBetween(start int, end int) Validator {
	return &lengthBetween{
		errorMsg: fmt.Sprintf("长度必须在 %d-%d 之间", start, end),
		start:    start,
		end:      end,
	}
}

func (l *lengthBetween) validate(value interface{}) bool {
	length := len(fmt.Sprintf("%v", value))
	return length >= l.start && length <= l.end
}

func (l *lengthBetween) ErrorMsg(msg string) Validator {
	l.errorMsg = msg
	return l
}

func (l *lengthBetween) getErrorMsg() string {
	return l.errorMsg
}
