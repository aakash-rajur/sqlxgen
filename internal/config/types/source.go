package types

import (
	"fmt"
	"strings"
)

type Source struct {
	Models  *Model `json:"models" yaml:"models"`
	Queries *Query `json:"queries" yaml:"queries"`
}

func (s *Source) String() string {
	if s == nil {
		return "Source{nil}"
	}

	content := strings.Join(
		[]string{
			fmt.Sprintf("models: %v", s.Models),
			fmt.Sprintf("queries: %v", s.Queries),
		},
		", ",
	)

	return fmt.Sprintf("Source{%s}", content)
}

func (s *Source) Merge(other *Source) *Source {
	if other == nil {
		return s
	}

	if s == nil {
		return other
	}

	s.Models = s.Models.Merge(other.Models)
	s.Queries = s.Queries.Merge(other.Queries)

	return s
}
