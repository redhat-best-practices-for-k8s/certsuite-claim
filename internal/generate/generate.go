// Copyright (C) 2020 Red Hat, Inc.
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
	"fmt"
	"github.com/a-h/generate"
	"io"
	"os"
)

// New sets up the generate.Generator and output file for generating a Claim in GoLang.  It does not apply any override
// settings.
func New(outputGoFile string, inputFiles ...string) (*generate.Generator, io.Writer, error) {
	schemas, err := generate.ReadInputFiles(inputFiles, false)
	if err != nil {
		return nil, nil, err
	}
	g := generate.New(schemas...)
	err = g.CreateTypes()
	if err != nil {
		return nil, nil, err
	}

	var w io.Writer = os.Stdout
	w, err = os.Create(outputGoFile)
	if err != nil {
		return nil, nil, err
	}
	return g, w, nil
}

// applyOverrideConfiguration applies a single override configuration.
func applyOverrideConfiguration(g *generate.Generator, property, parent, overrideType string) error {
	if _, ok := g.Structs[property]; ok {
		delete(g.Structs, property)
		if parentStruct, ok := g.Structs[parent]; ok {
			if replacementStruct, ok := parentStruct.Fields[property]; ok {
				replacementStruct.Type = overrideType
				parentStruct.Fields[property] = replacementStruct
				return nil
			}
			return fmt.Errorf("parent %s has no property: %s", parent, property)
		}
		return fmt.Errorf("no such parent: %s", parent)
	}
	return fmt.Errorf("no such property: %s", property)
}

// ApplyOverrideConfiguration applies all override configurations specified.
func ApplyOverrideConfiguration(g *generate.Generator, config *OverrideConfig) error {
	for key, value := range config.Properties {
		err := applyOverrideConfiguration(g, key, value.Parent, value.OverrideType)
		if err != nil {
			return err
		}
	}
	return nil
}
