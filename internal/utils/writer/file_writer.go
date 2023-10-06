package writer

import (
	"os"

	"github.com/joomcode/errorx"
)

type FileWriter struct {
	FullPath string
	Content  string
}

func (f *FileWriter) Write() error {
	err := os.Remove(f.FullPath)

	if err != nil && !os.IsNotExist(err) {
		return errorx.IllegalState.Wrap(err, "unable to remove existing generated query file")
	}

	return os.WriteFile(f.FullPath, []byte(f.Content), 0644)
}

func NewFileWriter(fullPath string, content string) Writer {
	return &FileWriter{
		FullPath: fullPath,
		Content:  content,
	}
}
