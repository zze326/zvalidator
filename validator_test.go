/**
 * @Author: zze
 * @Date: 2022/4/29 14:59
 * @Desc:
 */
package zvalidator

import (
	"fmt"
	"testing"
)

func Test1_1(t *testing.T) {
	name := "zze"
	age := 15
	errs := Errors{Validate(name, Required().ErrorMsg("名字不能为空")), Validate(age, GreaterThan(18).ErrorMsg("年龄必须大于18"))}
	if errs.HasError() {
		fmt.Println(errs.GetString())
	}
}
