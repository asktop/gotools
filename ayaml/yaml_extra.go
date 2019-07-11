package ayaml

import "io/ioutil"

//读取yaml文件
func ReadYaml(yamlPath string, out interface{}) (err error) {
	configBytes, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}
	return Unmarshal(configBytes, out)
}
