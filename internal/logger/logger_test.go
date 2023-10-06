package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	t.Run("NewLogger", func(t *testing.T) {
		content := &bytes.Buffer{}

		writer := bufio.NewWriter(content)

		args := &Args{
			LogLevel: "info",
			Format:   "text",
			Writer:   writer,
		}

		logger := NewLogger(args)

		if logger == nil {
			t.Error("want logger to be instantiated")
		}

		logger.Info("Hello World")

		err := writer.Flush()

		if err != nil {
			t.Error(err)
		}

		assert.Containsf(t, content.String(), `level=INFO msg="Hello World"`, "want content to contain 'Hello World' but got '%s'", content.String())
	})

	t.Run("NewLogger with JSON format", func(t *testing.T) {
		content := &bytes.Buffer{}

		writer := bufio.NewWriter(content)

		args := &Args{
			LogLevel: "info",
			Format:   "json",
			Writer:   writer,
		}

		logger := NewLogger(args)

		if logger == nil {
			t.Error("want logger to be instantiated")
		}

		logger.Info("Hello World")

		err := writer.Flush()

		if err != nil {
			t.Error(err)
		}

		want := make(map[string]interface{})

		err = json.Unmarshal(content.Bytes(), &want)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "INFO", want["level"])

		assert.Equal(t, "Hello World", want["msg"])
	})
}
