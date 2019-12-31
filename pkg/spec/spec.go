package spec

import (
	"bytes"
	"io"
	"log"
	"text/template"

	"gopkg.in/yaml.v3"
)

func NodeKindString(kind yaml.Kind) string {
	kinds := map[yaml.Kind]string{
		yaml.DocumentNode: "document",
		yaml.SequenceNode: "sequence",
		yaml.MappingNode:  "mapping",
		yaml.ScalarNode:   "scalar",
		yaml.AliasNode:    "alias",
	}
	return kinds[kind]
}

type RootNode yaml.Node

func PParse(indent int, node yaml.Node) error {
	if indent == 0 {
		log.Printf("|")
	}
	indentstr := ""
	for indentcount := 0; indentcount < indent; indentcount++ {
		indentstr += "  "
	}
	log.Printf("%s*kind<tag>: %v<%v>", indentstr, NodeKindString(node.Kind), node.Tag)
	log.Printf("%s val '%v'\n", indentstr, node.Value)
	if node.Content == nil {
		return nil
	}
	log.Printf("%s content:\n", indentstr)
	for _, subnode := range node.Content {
		PParse(indent+1, *subnode)
	}
	return nil
}

type Mod struct {
	Name  string
	Attrs map[string]interface{}
}

type HabiSpec struct {
	Vars map[string]string `yaml:".vars"`
	Mods map[string]Mod    `yaml:"mods"`
}

func Render(tmpl string, vars interface{}, e *error) string {
	if *e != nil {
		return ""
	}
	rendered := &bytes.Buffer{}
	t, err := template.New("text").Parse(tmpl)
	if err != nil {
		*e = err
		return ""
	}
	if err := t.Execute(rendered, vars); err != nil {
		*e = err
		return ""
	}
	return rendered.String()
}

func Parse(specreader io.Reader) (habispec HabiSpec, err error) {
	// var rawspec RawHabiSpec
	node := yaml.Node{}
	err = yaml.NewDecoder(specreader).Decode(&node)
	if err != nil {
		return habispec, err
	}
	var doc Document
	doc.Parse(node)
	// PParse(0, node)
	// pretty.YAML(os.Stdout, node)
	// log.Printf("node \n%v", node)
	return habispec, err

	// filesspec := []file.Spec{}
	// for _, rawfile := range rawspec.Files {
	// 	mode, err := file.ModeConv(Render(rawfile.Mode, rawspec.Vars, &err))
	// 	if err != nil {
	// 		return habispec, err
	// 	}
	// 	content := Render(rawfile.Content, rawspec.Vars, &err)
	// 	var sha256sum string
	// 	if rawfile.Sha256 == "" && rawfile.Src == "" {
	// 		sha256sum = fmt.Sprintf("%x", sha256.Sum256([]byte(content)))
	// 	}
	// 	if rawfile.Sha256 != "" {
	// 		sha256sum = Render(rawfile.Sha256, rawspec.Vars, &err)
	// 	}

	// 	filespec := file.Spec{
	// 		Src:     Render(rawfile.Src, rawspec.Vars, &err),
	// 		Dst:     Render(rawfile.Dst, rawspec.Vars, &err),
	// 		Sha256:  sha256sum,
	// 		SrcType: Render(rawfile.SrcType, rawspec.Vars, &err),
	// 		Content: content,
	// 		Mode:    mode,
	// 	}
	// 	filesspec = append(filesspec, filespec)
	// }
	// habispec = HabiSpec{
	// 	Cmd:   rawspec.Cmd,
	// 	Vars:  rawspec.Vars,
	// 	Files: filesspec,
	// }
	// return habispec, err
}
