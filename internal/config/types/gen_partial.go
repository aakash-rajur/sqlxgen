package types

import (
	"fmt"
)

type GenPartial struct {
	Path string `json:"path" yaml:"path"`
}

func (g *GenPartial) String() string {
	if g == nil {
		return "GenPartial{nil}"
	}

	return fmt.Sprintf("GenPartial{path: %s}", g.Path)
}

func (g *GenPartial) Merge(other *GenPartial) *GenPartial {
	if other == nil {
		return g
	}

	if g == nil {
		return other
	}

	if other.Path != "" {
		g.Path = other.Path
	}

	return g
}
