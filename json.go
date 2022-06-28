/**
 * @Author: zze
 * @Date: 2022/4/29 15:51
 * @Desc:
 */
package zvalidator

import "encoding/json"

type zjson struct {
	errorMsg string
}

// Json 格式校验
func Json() Validator {
	return &zjson{
		errorMsg: "不符合 Json 格式",
	}
}

func (j *zjson) validate(value interface{}) bool {
	v, isOk := value.(string)
	if isOk {
		var tmpMap map[string]interface{}
		err := json.Unmarshal([]byte(v), &tmpMap)
		if err != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

func (j *zjson) ErrorMsg(msg string) Validator {
	j.errorMsg = msg
	return j
}

func (j *zjson) getErrorMsg() string {
	return j.errorMsg
}
