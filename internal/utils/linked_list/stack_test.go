package linked_list

import (
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	t.Run("NewStack[int]()", func(t *testing.T) {
		stack := NewStack[int]()

		assert.NotNil(t, stack)

		assert.Nil(t, stack.top)

		assert.Equal(t, 0, stack.Length())

		assert.True(t, stack.IsEmpty())

		assert.NotNil(t, stack.ToSlice())

		assert.Equal(t, []int{}, stack.ToSlice())
	})

	t.Run("NewStack[string]()", func(t *testing.T) {
		stack := NewStack[string]()

		assert.NotNil(t, stack)

		assert.Nil(t, stack.top)

		assert.Equal(t, 0, stack.Length())

		assert.True(t, stack.IsEmpty())

		assert.NotNil(t, stack.ToSlice())

		assert.Equal(t, []string{}, stack.ToSlice())
	})
}

func TestFromSlice(t *testing.T) {
	t.Run("FromSlice[int]()", func(t *testing.T) {
		slice := []int{1, 2, 3}

		stack := FromSlice[int](slice)

		assert.NotNil(t, stack)

		assert.NotNil(t, stack.top)

		assert.Equal(t, 3, stack.Length())

		assert.False(t, stack.IsEmpty())

		assert.Equal(t, []int{3, 2, 1}, stack.ToSlice())
	})

	t.Run("FromSlice[string]()", func(t *testing.T) {
		slice := []string{"a", "b", "c"}

		stack := FromSlice[string](slice)

		assert.NotNil(t, stack)

		assert.NotNil(t, stack.top)

		assert.Equal(t, 3, stack.Length())

		assert.False(t, stack.IsEmpty())

		assert.Equal(t, []string{"c", "b", "a"}, stack.ToSlice())
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("Push[int]()", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(utils.PointerTo(1))

		assert.NotNil(t, stack.top)

		assert.Equal(t, 1, stack.Length())

		assert.False(t, stack.IsEmpty())

		assert.Equal(t, []int{1}, stack.ToSlice())
	})

	t.Run("Push[string]()", func(t *testing.T) {
		stack := NewStack[string]()

		stack.Push(utils.PointerTo("a"))

		assert.NotNil(t, stack.top)

		assert.Equal(t, 1, stack.Length())

		assert.False(t, stack.IsEmpty())

		assert.Equal(t, []string{"a"}, stack.ToSlice())
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("Pop[int]()", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(utils.PointerTo(1))
		stack.Push(utils.PointerTo(2))
		stack.Push(utils.PointerTo(3))

		value, notEmpty := stack.Pop()

		assert.True(t, notEmpty)

		assert.Equal(t, 3, *value)

		assert.NotNil(t, stack.top)

		assert.Equal(t, 2, stack.Length())

		assert.False(t, stack.IsEmpty())

		assert.Equal(t, []int{2, 1}, stack.ToSlice())
	})

	t.Run("Pop[string]()", func(t *testing.T) {
		stack := NewStack[string]()

		stack.Push(utils.PointerTo("a"))
		stack.Push(utils.PointerTo("b"))
		stack.Push(utils.PointerTo("c"))

		value, notEmpty := stack.Pop()

		assert.True(t, notEmpty)

		assert.Equal(t, "c", *value)

		assert.NotNil(t, stack.top)

		assert.Equal(t, 2, stack.Length())

		assert.False(t, stack.IsEmpty())

		assert.Equal(t, []string{"b", "a"}, stack.ToSlice())
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("Peek[int]()", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(utils.PointerTo(1))
		stack.Push(utils.PointerTo(2))
		stack.Push(utils.PointerTo(3))

		value, notEmpty := stack.Peek()

		assert.True(t, notEmpty)

		assert.Equal(t, 3, *value)

		assert.NotNil(t, stack.top)

		assert.Equal(t, 3, stack.Length())

		assert.False(t, stack.IsEmpty())

		assert.Equal(t, []int{3, 2, 1}, stack.ToSlice())
	})

	t.Run("Peek[string]()", func(t *testing.T) {
		stack := NewStack[string]()

		stack.Push(utils.PointerTo("a"))
		stack.Push(utils.PointerTo("b"))
		stack.Push(utils.PointerTo("c"))

		value, notEmpty := stack.Peek()

		assert.True(t, notEmpty)

		assert.Equal(t, "c", *value)

		assert.NotNil(t, stack.top)

		assert.Equal(t, 3, stack.Length())

		assert.False(t, stack.IsEmpty())

		assert.Equal(t, []string{"c", "b", "a"}, stack.ToSlice())
	})
}

func TestStack_Clear(t *testing.T) {
	t.Run("Clear[int]()", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(utils.PointerTo(1))
		stack.Push(utils.PointerTo(2))
		stack.Push(utils.PointerTo(3))

		stack.Clear()

		assert.Nil(t, stack.top)

		assert.Equal(t, 0, stack.Length())

		assert.True(t, stack.IsEmpty())

		assert.Equal(t, []int{}, stack.ToSlice())
	})

	t.Run("Clear[string]()", func(t *testing.T) {
		stack := NewStack[string]()

		stack.Push(utils.PointerTo("a"))
		stack.Push(utils.PointerTo("b"))
		stack.Push(utils.PointerTo("c"))

		stack.Clear()

		assert.Nil(t, stack.top)

		assert.Equal(t, 0, stack.Length())

		assert.True(t, stack.IsEmpty())

		assert.Equal(t, []string{}, stack.ToSlice())
	})
}

func TestStack_ToSlice(t *testing.T) {
	t.Run("ToSlice[int]()", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(utils.PointerTo(1))
		stack.Push(utils.PointerTo(2))
		stack.Push(utils.PointerTo(3))

		assert.Equal(t, []int{3, 2, 1}, stack.ToSlice())
	})

	t.Run("ToSlice[string]()", func(t *testing.T) {
		stack := NewStack[string]()

		stack.Push(utils.PointerTo("a"))
		stack.Push(utils.PointerTo("b"))
		stack.Push(utils.PointerTo("c"))

		assert.Equal(t, []string{"c", "b", "a"}, stack.ToSlice())
	})
}
