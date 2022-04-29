/**
 * @Author: zze
 * @Date: 2022/4/29 17:09
 * @Desc:
 */
package zvalidator

type Errors []error

func (errLst Errors) GetString() string {
	for _, err := range errLst {
		if err != nil {
			return err.Error()
		}
	}
	return ""
}

func (errLst Errors) GetError() error {
	for _, err := range errLst {
		if err != nil {
			return err
		}
	}
	return nil
}

func (errLst Errors) HasError() bool {
	for _, err := range errLst {
		if err != nil {
			return true
		}
	}
	return false
}
