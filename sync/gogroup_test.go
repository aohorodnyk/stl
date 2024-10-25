package sync_test

import (
	"fmt"
	"time"

	"github.com/aohorodnyk/stl/sync"
)

func ExampleGoGroup() {
	// Creating a GoGroup.
	var gg sync.GoGroup

	// Define a number of goroutines we want to run.
	const goroutines = 5

	// Create a slice we will fill from different goroutines.
	result := make([]string, goroutines)

	// Run multiple goroutines.
	for idx := range goroutines {
		gg.Go(func() {
			// Sleep up to 1 second.
			time.Sleep(time.Duration(idx) * time.Millisecond)

			// Since we update only a number by index, we should have no issues with race condition.
			// Don't use it in production, it's just an example.
			result[idx] = fmt.Sprintf("Slept %d milliseconds", idx)
		})
	}

	gg.Wait()

	fmt.Println(result)

	// Output:
	// [Slept 0 milliseconds Slept 1 milliseconds Slept 2 milliseconds Slept 3 milliseconds Slept 4 milliseconds]
}
