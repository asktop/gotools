package ayaml

import (
	"testing"
)

//yaml反序列化
//安装 go get -u gopkg.in/yaml.v2
func TestYaml(t *testing.T) {
	type Config struct {
		Name   string `yaml:"name"`
		Age    int    `yaml:"age"`
		Spouse struct {
			Name string `yaml:"name"`
			Age  int    `yaml:"age"`
		} `yaml:"spouse"`
	}

	var config Config
	err := ReadYaml("yaml_test.yaml", &config)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(config.Name)
	t.Log(config.Spouse.Name)
	t.Log(config.Spouse.Age)
}
