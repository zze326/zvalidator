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

type zint struct {
	errorMsg string
}

// Int 整型数字
func Int() Validator {
	return &zint{
		errorMsg: "必须是整型数字",
	}
}

func (i *zint) validate(value interface{}) bool {
	_, err := strconv.Atoi(fmt.Sprintf("%v", value))
	return err == nil
}

func (i *zint) ErrorMsg(msg string) Validator {
	i.errorMsg = msg
	return i
}

func (i *zint) getErrorMsg() string {
	return i.errorMsg
}
