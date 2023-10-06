package writer

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryWriter_Write(t *testing.T) {
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

	mw := NewMemoryWriters()

	for i, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			writer := mw.Creator(testCase.fields.FullPath, testCase.fields.Content)

			err := writer.Write()

			if testCase.error != nil {
				assert.NotNil(t, err)

				errMsgLeft := testCase.error.Error()

				errMsgRight := err.Error()

				assert.Containsf(t, errMsgRight, errMsgLeft, "want error %s but got %s", errMsgLeft, errMsgRight)
			} else {
				assert.Nil(t, err)

				pen := mw.Writers[i]

				assert.Equal(t, testCase.fields.FullPath, pen.FullPath)

				assert.Equal(t, testCase.fields.Content, pen.Content)
			}
		})
	}
}
