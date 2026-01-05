// Copyright (C) 2020-2026 Red Hat, Inc.
//
// This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
// warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program; if not, write to the Free
// Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.

package generate

import (
	"encoding/json"
	"os"
	"strings"
	"unicode"
)

const (
	defaultOverrideConfigFile = "./override.json"
)

// OverrideConfig provides a means of specifying override values for code generation.
type OverrideConfig struct {
	// Properties provides a means of specifying override settings for a particular property.
	Properties []OverrideSettings `json:"properties,omitempty" yaml:"properties,omitempty"`
}

// OverrideSettings contains the override values for a particular property.
type OverrideSettings struct {
	// Parent is the JSON Schema parent for a given property.
	Path []string `json:"path" yaml:"path"`
	// OverrideType is the type the code generator should emit for a given property.  Normally, this is `map[string]interface{}`.
	OverrideType string `json:"overrideType" yaml:"overrideType"`
}

// GetOverrideConfig extracts the override settings for code generation.
func GetOverrideConfig(overrideConfigFile string) (*OverrideConfig, error) {
	if overrideConfigFile == "" {
		overrideConfigFile = defaultOverrideConfigFile
	}

	contents, err := os.ReadFile(overrideConfigFile)
	if err != nil {
		return nil, err
	}

	overrideConfig := &OverrideConfig{}
	err = json.Unmarshal(contents, overrideConfig)

	// Makes sure the first character is upper case, as this is how the generator handles properties.
	for _, overrideProperty := range overrideConfig.Properties {
		i := 0
		for _, pathElement := range overrideProperty.Path {
			runes := []rune(pathElement)
			if unicode.IsLower(runes[0]) {
				//nolint:staticcheck
				overrideProperty.Path[i] = strings.Title(pathElement)
			}
			i++
		}
	}

	return overrideConfig, err
}
