package yaml

import (
	"fmt"
	"io/ioutil"
	"testing"
)

//yaml反序列化
//安装 go get -u gopkg.in/yaml.v2
func TestYaml(t *testing.T) {
	type Config struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
		Spouse struct {
			Name string `yaml:"name"`
			Age  int    `yaml:"age"`
		} `yaml:"spouse"`
	}

	var config Config
	configFile := "yaml_test.yaml"
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	Unmarshal(configBytes, &config)

	fmt.Println(config.Name)
	fmt.Println(config.Spouse.Name)
	fmt.Println(config.Spouse.Age)
}
