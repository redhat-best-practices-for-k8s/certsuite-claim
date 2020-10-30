package claim_test

import (
	"encoding/json"
	"github.com/redhat-nfvpe/test-network-function-claim/pkg/claim"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"testing"
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
		assert.Equal(t, testCaseDefinition.expectedMarshallJSONError, err != nil)

		if testCaseDefinition.expectedMarshallJSONError == false {
			// start time assertion
			assert.Equal(t, "1970-01-01T10:05:08+01:00", root.Claim.StartTime)
			assert.Equal(t, "1970-01-01T10:05:08+01:00", root.Claim.EndTime)

			generatedContents, err := json.Marshal(root)
			assert.Nil(t, err)
			assert.NotNil(t, generatedContents)
		}
	}
}
