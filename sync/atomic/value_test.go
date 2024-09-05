package atomic_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/sync/atomic"
)

func ExampleValue() {
	type myCustomType struct {
		str string
	}

	var (
		value       myCustomType
		valueNew    myCustomType
		valueAtomic atomic.Value[myCustomType]
		valueLoaded *myCustomType
		swapped     bool
	)

	value.str = "My custom string"
	valueNew.str = "String from new value"

	valueLoaded = valueAtomic.Load()
	fmt.Printf("Loaded nil: %v\n", valueLoaded)

	valueAtomic.Store(&value)

	valueLoaded = valueAtomic.Load()
	fmt.Printf("Loaded value: %v\n", valueLoaded)

	valueSwapped := valueAtomic.Swap(&valueNew)
	fmt.Printf("Loaded swapped: %v\n", valueSwapped)

	valueLoaded = valueAtomic.Load()
	fmt.Printf("Loaded new: %v\n", valueLoaded)

	swapped = valueAtomic.CompareAndSwap(&value, &valueNew)
	fmt.Printf("Value not swapped: %v\n", swapped)

	valueLoaded = valueAtomic.Load()
	fmt.Printf("Loaded new: %v\n", valueLoaded)

	swapped = valueAtomic.CompareAndSwap(&valueNew, &value)
	fmt.Printf("Value swapped: %v\n", swapped)

	valueLoaded = valueAtomic.Load()
	fmt.Printf("Loaded new: %v\n", valueLoaded)

	// Output:
	// Loaded nil: <nil>
	// Loaded value: &{My custom string}
	// Loaded swapped: &{My custom string}
	// Loaded new: &{String from new value}
	// Value not swapped: false
	// Loaded new: &{String from new value}
	// Value swapped: true
	// Loaded new: &{My custom string}
}
