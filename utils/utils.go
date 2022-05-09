/**
 * @Author: zze
 * @Date: 2022/5/9 15:12
 * @Desc:
 */
package utils

import "gopkg.in/yaml.v2"

func UnmarshalYamlToMap(yamlStr string) (map[interface{}]interface{}, error) {
	var m map[interface{}]interface{}
	if err := yaml.Unmarshal([]byte(yamlStr), &m); err != nil {
		return nil, err
	}
	return m, nil
}
