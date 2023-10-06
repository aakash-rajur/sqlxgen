package types

import (
	"fmt"
	"strings"
)

type Query struct {
	Paths   []string `json:"paths" yaml:"paths"`
	Include []string `json:"include" yaml:"include"`
	Exclude []string `json:"exclude" yaml:"exclude"`
}

func (q *Query) String() string {
	if q == nil {
		return "Query{nil}"
	}

	content := strings.Join(
		[]string{
			fmt.Sprintf("paths: %v", q.Paths),
			fmt.Sprintf("include: %v", q.Include),
			fmt.Sprintf("exclude: %v", q.Exclude),
		},
		", ",
	)

	return fmt.Sprintf("Query{%s}", content)
}

func (q *Query) Merge(other *Query) *Query {
	if other == nil {
		return q
	}

	if q == nil {
		return other
	}

	if other.Paths != nil {
		q.Paths = other.Paths
	}

	if other.Include != nil {
		q.Include = other.Include
	}

	if other.Exclude != nil {
		q.Exclude = other.Exclude
	}

	return q
}
