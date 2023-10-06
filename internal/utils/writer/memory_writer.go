package writer

type MemoryWriter struct {
	FullPath string
	Content  string
	OnWrite  func(mw *MemoryWriter) error
}

func (m *MemoryWriter) Write() error {
	if m.OnWrite != nil {
		return m.OnWrite(m)
	}

	return nil
}

type MemoryWriters struct {
	Writers []*MemoryWriter
}

func (mws *MemoryWriters) Creator(fullPath string, content string) Writer {
	mw := &MemoryWriter{
		FullPath: fullPath,
		Content:  content,
		OnWrite: func(mw *MemoryWriter) error {
			mws.Writers = append(mws.Writers, mw)

			return nil
		},
	}

	return mw
}

func NewMemoryWriters() *MemoryWriters {
	return &MemoryWriters{
		Writers: make([]*MemoryWriter, 0),
	}
}
