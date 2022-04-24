package stack_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/stack"
)

var (
	_ stack.Stack[int]               = &stack.Sync[int]{}
	_ stack.Stack[string]            = &stack.Sync[string]{}
	_ stack.Stack[map[string]string] = &stack.Sync[map[string]string]{}
	_ stack.Stack[string]            = stack.NewSync[string](&stack.Dynamic[string]{})
	_ stack.Stack[string]            = stack.NewSync[string](&stack.Fixed[string]{})
)

func TestDynamicClearSync(t *testing.T) {
	t.Parallel()

	stackObj := stack.NewDynamicSync[int]()

	clearTest(t, stackObj)
}

func TestDynamicSize10Sync(t *testing.T) {
	t.Parallel()

	stackObj := stack.NewDynamicSync[string]()

	size10DynamicTest(t, stackObj)
}

func TestFixedSync(t *testing.T) {
	t.Parallel()

	stackObj := stack.NewFixedSync[int](10)

	size10Test(t, stackObj)
}

func TestFixedClearSync(t *testing.T) {
	t.Parallel()

	stackObj := stack.NewFixedSync[int](3)

	clearTest(t, stackObj)
}
