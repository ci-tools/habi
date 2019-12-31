package spec

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

// Document ...
type Document struct {
	Vars    map[string]string
	Mods    map[string]Mod
	Stanzas map[string]Stanza
}

// Parse ...
func (doc Document) Parse(docNode yaml.Node) (err error) {
	docContent := docNode.Content[0]
	if docContent.Kind != yaml.MappingNode {
		return fmt.Errorf("document should be map")
	}
	for key, val := range yamlMapping(docContent.Content) {
		var stanza Stanza
		for _, stanzaType := range stanzaTypes {
			// log.Printf("%s\n", key.Value)
			if stanzaType.Match(key, val) {
				stanza, err = stanzaType.Parse(key, val)
				if err != nil {
					return err
				}
				break
			}
		}
		if stanza == nil {
			// default action ?
			continue
		}
		log.Printf("\n%v", stanza.String())
	}
	return nil
}

type stanzaType struct {
	Match func(key *yaml.Node, val *yaml.Node) bool
	Parse func(key *yaml.Node, val *yaml.Node) (Stanza, error)
}

var stanzaTypes = []stanzaType{
	stanzaVarsType,
	stanzaModType,
}

// Stanza ...
type Stanza interface {
	String() string
}

// StanzaVars ...
type StanzaVars map[string]string

// String ...
func (vars StanzaVars) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, ".vars:\n")
	for varKey, varVal := range vars {
		fmt.Fprintf(&sb, "\t%s: %s\n", varKey, varVal)
	}
	return sb.String()
}

var stanzaVarsType = stanzaType{
	Match: func(key *yaml.Node, val *yaml.Node) bool {
		return key.Value == ".vars"
	},
	Parse: func(key *yaml.Node, val *yaml.Node) (Stanza, error) {
		vars := StanzaVars{}
		_, vars = formatYamlMap(key, val)
		return vars, nil
	},
}

var stanzaModType = stanzaType{
	Match: func(key *yaml.Node, val *yaml.Node) bool {
		return string(key.Value[0]) != "." && val.Kind == yaml.SequenceNode
	},
	Parse: func(key *yaml.Node, val *yaml.Node) (Stanza, error) {
		mod := StanzaMod{
			Name:   key.Value,
			Claims: []Claim{},
		}
		for _, claimNode := range val.Content {
			mod.Claims = append(mod.Claims, ClaimParse(claimNode))
		}
		return mod, nil
	},
}

// KeyWordsAndPrefixes ...
var KeyWordsAndPrefixes = []string{
	"as", "for", "if",
}

// IsKeyWord check if string matches any keyword
func IsKeyWord(str string) bool {
	for _, keyWord := range KeyWordsAndPrefixes {
		keyWordPrefix := fmt.Sprintf("%s-", keyWord)
		if str == keyWord || strings.HasPrefix(str, keyWordPrefix) {
			return true
		}
	}
	return false
}

// ClaimControl ...
type ClaimControl struct {
	Name   string
	Params map[string]string
}

func (cc ClaimControl) String() string {
	sb := strings.Builder{}
	fmt.Fprintf(&sb, "  [%s]:\n", cc.Name)
	for paramName, paramVal := range cc.Params {
		fmt.Fprintf(&sb, "    %s: %s\n", paramName, paramVal)
	}
	return sb.String()
}

// Claim ...
type Claim struct {
	Name     string
	Params   map[string]string
	Controls map[string]ClaimControl
}

func (c Claim) String() string {
	sb := strings.Builder{}
	fmt.Fprintf(&sb, "  %s:\n", c.Name)
	for paramName, paramVal := range c.Params {
		fmt.Fprintf(&sb, "    %s: %s\n", paramName, paramVal)
	}
	for _, control := range c.Controls {
		fmt.Fprintf(&sb, "%s\n", control)
	}
	return sb.String()
}

// ClaimParse ...
func ClaimParse(claimNode *yaml.Node) Claim {
	claim := Claim{
		Params:   map[string]string{},
		Controls: map[string]ClaimControl{},
	}
	switch claimNode.Kind {
	case yaml.MappingNode:
		for claimKey, claimVal := range yamlMapping(claimNode.Content) {
			switch {
			case IsKeyWord(claimKey.Value):
				control := ClaimControl{
					Name:   claimKey.Value,
					Params: map[string]string{},
				}
				control.Name, control.Params = formatYamlMap(claimKey, claimVal)
				claim.Controls[control.Name] = control
				continue
			default:
				claim.Name, claim.Params = formatYamlMap(claimKey, claimVal)
			}
		}
	case yaml.ScalarNode:
		claim.Name = claimNode.Value
	}
	return claim
}

// StanzaMod ...
type StanzaMod struct {
	Name   string
	Claims []Claim
}

// String ...
func (mod StanzaMod) String() string {
	sb := strings.Builder{}
	fmt.Fprintf(&sb, "%s:\n", mod.Name)
	for _, claim := range mod.Claims {
		fmt.Fprintf(&sb, "\n%s\n", claim)
	}
	return sb.String()
}

// StanzaModFile ...
type StanzaModFile struct {
	Src    string
	Dst    string
	Sha256 string
	Mode   string
}

// StanzaModCmd ...
type StanzaModCmd string
