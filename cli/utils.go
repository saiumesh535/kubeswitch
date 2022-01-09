package cli

import "github.com/ghodss/yaml"

func YamlToJson(input []byte) ([]byte, error) {
	return yaml.YAMLToJSON(input)
}
