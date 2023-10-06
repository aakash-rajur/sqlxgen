package types

import (
	"fmt"
	"strings"
)

type Model struct {
	Schemas []string `json:"schemas" yaml:"schemas"`
	Include []string `json:"include" yaml:"include"`
	Exclude []string `json:"exclude" yaml:"exclude"`
}

func (m *Model) String() string {
	if m == nil {
		return "Model{nil}"
	}

	content := strings.Join(
		[]string{
			fmt.Sprintf("schemas: %v", m.Schemas),
			fmt.Sprintf("include: %v", m.Include),
			fmt.Sprintf("exclude: %v", m.Exclude),
		},
		", ",
	)

	return fmt.Sprintf("Model{%s}", content)
}

func (m *Model) Merge(other *Model) *Model {
	if other == nil {
		return m
	}

	if m == nil {
		return other
	}

	if other.Schemas != nil {
		m.Schemas = other.Schemas
	}

	if other.Include != nil {
		m.Include = other.Include
	}

	if other.Exclude != nil {
		m.Exclude = other.Exclude
	}

	return m
}
