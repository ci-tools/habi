package pretty

import (
	"io"

	"gopkg.in/yaml.v3"
)

func YAML(w io.Writer, yamlData interface{}) {
	encoder := yaml.NewEncoder(w)
	encoder.Encode(yamlData)
}
