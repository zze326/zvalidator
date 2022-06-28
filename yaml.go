/**
 * @Author: zze
 * @Date: 2022/4/29 15:51
 * @Desc:
 */
package zvalidator

import "github.com/zze326/zvalidator/utils"

type yaml struct {
	errorMsg string
}

// Yaml 格式校验
func Yaml() Validator {
	return &yaml{
		errorMsg: "不符合 Yaml 格式",
	}
}

func (y *yaml) validate(value interface{}) bool {
	v, isOk := value.(string)
	if isOk {
		_, err := utils.UnmarshalYamlToMap(v)
		if err != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

func (y *yaml) ErrorMsg(msg string) Validator {
	y.errorMsg = msg
	return y
}

func (y *yaml) getErrorMsg() string {
	return y.errorMsg
}
