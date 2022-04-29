/**
 * @Author: zze
 * @Date: 2022/4/29 16:47
 * @Desc:
 */
package zvalidator

import (
	"fmt"
	"strconv"
)

type between struct {
	errorMsg string
	start    float64
	end      float64
}

// Between 值范围
func Between(start float64, end float64) Validator {
	return &between{
		errorMsg: fmt.Sprintf("必须在 %v-%v 之间", start, end),
		start:    start,
		end:      end,
	}
}

func (b *between) validate(value interface{}) bool {
	result, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	if err != nil {
		b.errorMsg = fmt.Sprintf("%v 不是一个数字", value)
		return false
	}
	return result >= b.start && result <= b.end
}

func (b *between) ErrorMsg(msg string) Validator {
	b.errorMsg = msg
	return b
}

func (b *between) getErrorMsg() string {
	return b.errorMsg
}
