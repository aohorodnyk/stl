package queue_test

import (
	"testing"

	"github.com/aohorodnyk/stl/collections/queue"
)

var (
	_ queue.Queue[int]    = &queue.Sync[int]{}
	_ queue.Queue[string] = queue.NewSync[string](&queue.Dynamic[string]{})
	_ queue.Queue[string] = queue.NewSync[string](&queue.Fixed[string]{})
)

func TestDynamicSyncSize10(t *testing.T) {
	t.Parallel()

	queueObj := queue.NewDynamicSync[string]()

	size10DynamicTest(t, queueObj)
}

func TestDynamicSyncClear(t *testing.T) {
	t.Parallel()

	queueObj := queue.NewDynamicSync[int]()

	clearTest(t, queueObj)
}

func TestFixedSync(t *testing.T) {
	t.Parallel()

	queueObj := queue.NewFixedSync[int](10)

	size10Test(t, queueObj)
}

func TestFixedSyncClear(t *testing.T) {
	t.Parallel()

	queueObj := queue.NewFixedSync[int](10)

	size10Test(t, queueObj)
}
