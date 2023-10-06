package types

import (
	"fmt"
	"strings"
)

type Gen struct {
	Store *GenPartial `json:"store" yaml:"store"`
	Model *GenPartial `json:"models" yaml:"models"`
}

func (g *Gen) String() string {
	if g == nil {
		return "Gen{nil}"
	}

	content := strings.Join(
		[]string{
			fmt.Sprintf("store: %v", g.Store),
			fmt.Sprintf("model: %v", g.Model),
		},
		", ",
	)

	return fmt.Sprintf("Gen{%s}", content)
}

func (g *Gen) Merge(other *Gen) *Gen {
	if other == nil {
		return g
	}

	if g == nil {
		return other
	}

	g.Store = g.Store.Merge(other.Store)
	g.Model = g.Model.Merge(other.Model)

	return g
}
