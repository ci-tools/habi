package spec

import "gopkg.in/yaml.v3"

func yamlMapping(nodes []*yaml.Node) map[*yaml.Node]*yaml.Node {
	nodesMap := map[*yaml.Node]*yaml.Node{}
	for keyIndex := 0; keyIndex < len(nodes); keyIndex += 2 {
		valIndex := keyIndex + 1
		nodesMap[nodes[keyIndex]] = nodes[valIndex]
	}
	return nodesMap
}

func formatYamlMap(key *yaml.Node, val *yaml.Node) (string, map[string]string) {
	name := key.Value
	formatedMap := map[string]string{}
	switch val.Kind {
	case yaml.MappingNode:
		for paramKey, paramVal := range yamlMapping(val.Content) {
			formatedMap[paramKey.Value] = paramVal.Value
		}
	case yaml.ScalarNode:
		formatedMap[".short-form"] = val.Value
	}
	return name, formatedMap
}
