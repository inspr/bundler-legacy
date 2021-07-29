package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var defaultPath string
var inputPath string

type TestYaml struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func isYaml(file string) bool {
	end := filepath.Ext(file)
	if end != ".yml" && end != ".yaml" {

		return false
	}
	return true
}

func validFile(filePath string) bool {
	var err error
	if _, err = os.Stat(filePath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

func readFile() {
	var test TestYaml
	yamlFile, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fmt.Println(err)
	}
	yaml.Unmarshal(yamlFile, &test)
	fmt.Println(test)
}
