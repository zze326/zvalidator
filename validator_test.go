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

func Test1_2(t *testing.T) {
	numStr := "123"
	var num int
	err := Validate(numStr, IntReturn(&num))
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}

func Test1_3(t *testing.T) {
	yamlStr := "a: 12"
	err := Validate(yamlStr, Yaml())
	if err != nil {
		panic(err)
	}
}
