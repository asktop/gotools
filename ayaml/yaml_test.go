package ayaml

import (
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
		t.Log(err)
		return
	}
	Unmarshal(configBytes, &config)

	t.Log(config.Name)
	t.Log(config.Spouse.Name)
	t.Log(config.Spouse.Age)
}
