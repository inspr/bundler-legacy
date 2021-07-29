package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var Path string
var fpath string

type TestYaml struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func isYaml() bool {
	end := filepath.Ext(fpath)
	if end != ".yml" && end != ".yaml" {

		return false
	}
	return true
}

func readFile() {
	var test TestYaml
	yamlFile, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println(err)
	}
	yaml.Unmarshal(yamlFile, &test)
	fmt.Println(test)
}
