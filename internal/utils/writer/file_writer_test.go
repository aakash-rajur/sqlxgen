package writer

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileWriter_Write(t *testing.T) {
	dir := t.TempDir()

	type fields struct {
		FullPath string
		Content  string
	}

	testCases := []struct {
		testName string
		fields   fields
		error    error
	}{
		{
			testName: "valid file write 1",
			fields: fields{
				FullPath: path.Join(dir, "test1.txt"),
				Content:  "lorem ipsum dolor sit amet",
			},
			error: nil,
		},
		{
			testName: "valid file write 2",
			fields: fields{
				FullPath: path.Join(dir, "test2.txt"),
				Content:  "lorem ipsum dolor sit amet",
			},
			error: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			writer := NewFileWriter(testCase.fields.FullPath, testCase.fields.Content)

			err := writer.Write()

			if testCase.error != nil {
				assert.NotNil(t, err)

				errMsgLeft := testCase.error.Error()

				errMsgRight := err.Error()

				assert.Containsf(t, errMsgRight, errMsgLeft, "want error %s but got %s", errMsgLeft, errMsgRight)
			} else {
				assert.Nil(t, err)

				content, err := os.ReadFile(testCase.fields.FullPath)

				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(t, testCase.fields.Content, string(content), "want content %s but got %s", testCase.fields.Content, string(content))
			}
		})
	}
}
