# zvalidator

```bash
go get github.com/zze326/zvalidator
```

# Usage

```go
package t_zvalidator_test

import (
	"fmt"
	"github.com/zze326/zvalidator"
	"testing"
)

func Test_1(t *testing.T) {
	name := "zze"
	age := 15
	errs := zvalidator.Errors{
		zvalidator.Validate(name, zvalidator.Required().ErrorMsg("名字不能为空")),
		zvalidator.Validate(age, zvalidator.GreaterThan(18).ErrorMsg("年龄必须大于 18"))}
	if errs.HasError() {
		// 如果是 Web 开发此处可以中断请求流程将错误信息响应给客户端
		fmt.Println(errs.GetString())
	}
}
```
