/**
 * @Author: zze
 * @Date: 2022/4/29 15:27
 * @Desc:
 */
package zvalidator

import "fmt"

type Validator interface {
	validate(value interface{}) bool
	getErrorMsg() string
	ErrorMsg(string) Validator
}

func Validate(value interface{}, validators ...Validator) error {
	for _, validator := range validators {
		isOk := validator.validate(value)
		if !isOk {
			return fmt.Errorf(validator.getErrorMsg())
		}
	}
	return nil
}
