package writer

type Writer interface {
	Write() error
}

type Creator func(fullPath string, content string) Writer
