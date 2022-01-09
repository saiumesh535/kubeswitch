package cli

import (
	"os"

	"github.com/tidwall/gjson"
)

const PATH = "/.kube/config"

type KubeConfig struct {
	Contexts       []string
	CurrentContext string
}

func GetConfigFromFile() ([]byte, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configPath := homeDir + PATH
	return os.ReadFile(configPath)
}

func GetKubeConfig() (KubeConfig, error) {
	config, err := GetConfigFromFile()
	if err != nil {
		return KubeConfig{}, err
	}
	// now convert YAML to json
	kubeConfig, err := YamlToJson(config)
	if err != nil {
		return KubeConfig{}, err
	}

	// now extract contexts from kubeconfig
	var contexts []string

	gjson.Get(string(kubeConfig), "contexts").ForEach(func(key, value gjson.Result) bool {
		currentItem := value.Get("name").String()
		contexts = append(contexts, currentItem)
		return true
	})

	currentContext := gjson.Get(string(kubeConfig), "current-context").String()

	return KubeConfig{
		Contexts:       contexts,
		CurrentContext: currentContext,
	}, nil

}
