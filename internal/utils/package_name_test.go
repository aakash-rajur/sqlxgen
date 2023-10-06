package utils

import (
	"os"
	"path"
	"testing"

	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectPackageName(t *testing.T) {
	testCases := []struct {
		testName    string
		content     string
		error       error
		packageName string
	}{
		{
			testName: "valid go.mod file1",
			content: `
module github.com/username/project

go 1.16
`,
			error:       nil,
			packageName: "github.com/username/project",
		},
		{
			testName: "valid go.mod file2",
			content: `
module github.com/aakash-rajur/sqlxgen

go 1.21.1

require (
	github.com/deckarep/golang-set v1.8.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/jinzhu/inflection v1.0.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/joho/godotenv v1.5.1
	github.com/joomcode/errorx v1.1.1
	github.com/lib/pq v1.10.9
	golang.org/x/mod v0.12.0
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/stretchr/testify v1.8.2 // indirect

`,
			error:       nil,
			packageName: "github.com/aakash-rajur/sqlxgen",
		},
		{
			testName: "invalid go.mod file",
			content: `
lorem ipsum dolor sit amet
`,
			error:       errorx.InitializationFailed.Wrap(nil, "unable to parse go.mod file"),
			packageName: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			dir := t.TempDir()

			filepath := path.Join(dir, "go.mod")

			err := os.WriteFile(filepath, []byte(testCase.content), 0644)

			if err != nil {
				t.Fatal(err)
			}

			packageName, err := GetProjectPackageName(dir)

			if testCase.error != nil {
				assert.NotNil(t, err, "want error but got nil")

				errMsgLeft := testCase.error.Error()

				errMsgRight := err.Error()

				assert.Containsf(t, errMsgRight, errMsgLeft, "want error %s but got %s", errMsgLeft, errMsgRight)
			} else {
				assert.Nil(t, err, "want no error but got %s", err)

				assert.Equal(t, testCase.packageName, packageName, "want package name %s but got %s", testCase.packageName, packageName)
			}
		})
	}
}
