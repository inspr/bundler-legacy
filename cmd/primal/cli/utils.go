package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"inspr.dev/primal/pkg/api"
)

var defaultPath string
var inputPath string

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

func getConfigs(path string) (api.PrimalOptions, error) {
	var config api.PrimalOptions

	bContent, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return api.PrimalOptions{}, err
	}

	yaml.Unmarshal(bContent, &config)

	return config, nil
}

func getDirPath(path string) string {
	dir := strings.Split(path, "/")
	dir = dir[:len(dir)-1]
	newDir := strings.Join(dir, "/")
	return newDir
}
