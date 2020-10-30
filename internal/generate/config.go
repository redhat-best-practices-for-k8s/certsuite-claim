package generate

import (
	"encoding/json"
	"io/ioutil"
)

const (
	defaultOverrideConfigFile = "./override.json"
)

// OverrideConfig provides a means of specifying override values for code generation.
type OverrideConfig struct {
	// Properties provides a means of specifying override settings for a particular property.
	Properties map[string]OverrideSettings `json:"properties,omitempty" yaml:"properties,omitempty"`
}

// OverrideSettings contains the override values for a particular property.
type OverrideSettings struct {
	// Parent is the JSON Schema parent for a given property.
	Parent string `json:"parent" yaml:"parent"`
	// OverrideType is the type the code generator should emit for a given property.  Normally, this is `map[string]interface{}`.
	OverrideType string `json:"overrideType" yaml:"overrideType"`
}

// GetOverrideConfig extracts the override settings for code generation.
func GetOverrideConfig(overrideConfigFile string) (*OverrideConfig, error) {
	if overrideConfigFile == "" {
		overrideConfigFile = defaultOverrideConfigFile
	}

	contents, err := ioutil.ReadFile(overrideConfigFile)
	if err != nil {
		return nil, err
	}

	overrideConfig := &OverrideConfig{}
	err = json.Unmarshal(contents, overrideConfig)
	return overrideConfig, err
}
