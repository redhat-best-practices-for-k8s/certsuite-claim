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

package claim_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/redhat-nfvpe/test-network-function-claim/pkg/claim"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	expectedMarshallJSONError bool
	expectedStartTime         string
	expectedEndTime           string
}

var testCases = map[string]*testCase{
	"claim-valid": {
		expectedMarshallJSONError: false,
		expectedStartTime:         "1970-01-01T10:05:08+01:00",
		expectedEndTime:           "1970-01-01T10:05:08+01:00",
	},
	"claim-invalid-junit-payload": {
		expectedMarshallJSONError: true,
	},
	"claim-invalid-additional-property": {
		expectedMarshallJSONError: true,
	},
	"invalid-json": {
		expectedMarshallJSONError: true,
	},
	"missing-claim": {
		expectedMarshallJSONError: true,
	},
}

func getTestFile(testCaseName string) string {
	return path.Join("testdata", testCaseName+".json")
}

func getTestFileContents(testCaseName string) ([]byte, error) {
	testFilePath := getTestFile(testCaseName)
	return ioutil.ReadFile(testFilePath)
}

func TestRoot_MarshalJSON(t *testing.T) {
	for testCaseName, testCaseDefinition := range testCases {
		contents, err := getTestFileContents(testCaseName)

		// raw data read tests
		assert.Nil(t, err)
		assert.NotNil(t, contents)

		// try to UnmarshallJSON the input
		root := &claim.Root{}
		err = json.Unmarshal(contents, root)
		fmt.Println(testCaseName)
		assert.Equal(t, testCaseDefinition.expectedMarshallJSONError, err != nil)

		if testCaseDefinition.expectedMarshallJSONError == false {
			// start time assertion
			assert.Equal(t, "1970-01-01T10:05:08+01:00", root.Claim.Metadata.StartTime)
			assert.Equal(t, "1970-01-01T10:05:08+01:00", root.Claim.Metadata.EndTime)

			generatedContents, err := json.Marshal(root)
			assert.Nil(t, err)
			assert.NotNil(t, generatedContents)
		}
	}
}
