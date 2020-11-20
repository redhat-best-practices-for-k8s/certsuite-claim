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
	"io"
	"os"
	"regexp"

	"github.com/a-h/generate"
)

const (
	fieldTypeRegexIndex  = 1
	fieldTypeRegexString = `[^\*]*\*(\w+)`
)

var fieldTypeRegex = regexp.MustCompile(fieldTypeRegexString)

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

	var w io.Writer
	w, err = os.Create(outputGoFile)
	if err != nil {
		return nil, nil, err
	}
	return g, w, nil
}

func apply(parentStruct *generate.Struct, property, overrideType string, path []string) error {
	if replacementField, ok := parentStruct.Fields[property]; ok {
		replacementField.Type = overrideType
		parentStruct.Fields[property] = replacementField
		return nil
	}
	return fmt.Errorf("incorrect path: %s", path)
}

func getFieldType(typ string) (string, error) {
	matches := fieldTypeRegex.FindStringSubmatch(typ)
	if len(matches) >= fieldTypeRegexIndex {
		return matches[fieldTypeRegexIndex], nil
	}
	return "", fmt.Errorf("couldn't determine type for erasure: %s", typ)
}

// applyOverrideConfiguration applies a single override configuration.
func applyOverrideConfiguration(g *generate.Generator, overrideType string, path []string) error {
	pathLength := len(path)
	if pathLength <= 1 {
		return fmt.Errorf("path to replacement field must have at least two elements: %s", path)
	}
	structName := path[pathLength-2]
	if parentStruct, ok := g.Structs[structName]; ok {
		fieldName := path[pathLength-1]
		if field, ok := parentStruct.Fields[fieldName]; ok {
			fieldType, err := getFieldType(field.Type)
			if err != nil {
				return err
			}
			delete(g.Structs, fieldType)
			return apply(&parentStruct, fieldName, overrideType, path)
		}
	}
	return fmt.Errorf("incorrect path: %s", path)
}

// ApplyOverrideConfiguration applies all override configurations specified.
func ApplyOverrideConfiguration(g *generate.Generator, config *OverrideConfig) error {
	for _, value := range config.Properties {
		err := applyOverrideConfiguration(g, value.OverrideType, value.Path)
		if err != nil {
			return err
		}
	}
	return nil
}
