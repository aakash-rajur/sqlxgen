package prepare

import (
	"regexp"

	"github.com/aakash-rajur/sqlxgen/internal/utils/linked_list"
	"github.com/joomcode/errorx"
)

func splitWhereContexts(query string) ([]string, error) {
	if query == "" {
		return []string{}, nil
	}

	partials := make([]string, 0)

	matches, err := findAllWhereMatches(query)

	if err != nil {
		return partials, errorx.IllegalFormat.Wrap(err, "failed to find where matches")
	}

	if len(matches) == 0 {
		return []string{query}, nil
	}

	parens := findAllParentheses(query)

	for _, match := range matches {
		whereStart, whereEnd := match[4], match[5]-1

		closestParen := parentheses{
			Start: -1,
			End:   len(query),
		}

		for _, paren := range parens {
			isWithin := paren.Start < whereStart && whereEnd < paren.End

			if !isWithin {
				continue
			}

			isCloser := closestParen.Start < paren.Start && paren.End < closestParen.End

			if !isCloser {
				continue
			}

			closestParen = paren
		}

		if closestParen.Start == -1 {
			continue
		}

		wherePartial := query[whereStart:closestParen.End]

		if wherePartial == "" {
			continue
		}

		partials = append(partials, wherePartial)
	}

	if len(partials) == 0 {
		return []string{query}, nil
	}

	return partials, nil
}

func findAllWhereMatches(query string) ([][]int, error) {
	re, err := regexp.Compile(`(\)|\s)(where)(\(|\s)`)

	if err != nil {
		return [][]int{}, errorx.IllegalFormat.Wrap(err, "failed to compile with regex")
	}

	matches := re.FindAllStringSubmatchIndex(query, -1)

	if len(matches) == 0 {
		return [][]int{}, nil
	}

	return matches, nil
}

func findAllParentheses(content string) []parentheses {
	ps := make([]parentheses, 0)

	stack := linked_list.NewStack[parenthesesNode]()

	for index, character := range content {
		if character == '(' {
			node := parenthesesNode{
				Character: character,
				Index:     index,
			}

			stack.Push(&node)
		}

		if character == ')' {
			prevNode, _ := stack.Peek()

			if prevNode.Character != '(' {
				continue
			}

			stack.Pop()

			p := parentheses{
				Start: prevNode.Index,
				End:   index,
			}

			ps = append(ps, p)
		}
	}

	return ps
}

type parentheses struct {
	Start int
	End   int
}

type parenthesesNode struct {
	Character rune
	Index     int
}
