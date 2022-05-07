package collections_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/collections"
)

func ExampleCounter() {
	counter := collections.NewCounter[string]()
	counter.Add("a", 123)
	counter.Add("a", 43535)

	counter.Inc("b")
	counter.Inc("b")

	counter.Add("c", 1632)

	fmt.Println(counter.Counted("a"))
	fmt.Println(counter.Counted("b"))
	fmt.Println(counter.Counted("c"))

	counter.Dec("b")
	fmt.Println(counter.Counted("b"))

	counter.Sub("c", 42)
	fmt.Println(counter.Counted("c"))

	counter.Add("d", 234)
	fmt.Println(counter.Counted("d"))

	counter.Add("e", 7435)
	fmt.Println(counter.Counted("e"))

	counter.Add("f", 4361)
	fmt.Println(counter.Counted("f"))

	fmt.Println(counter.Sum())

	fmt.Println(counter.Max())
	fmt.Println(counter.Min())

	fmt.Println(counter.MaxAll())
	fmt.Println(counter.MinAll())

	fmt.Println(counter.ToSlice())

	fmt.Println(counter.Top(3))
	fmt.Println(counter.Bottom(3))

	fmt.Println(counter.Length())

	fmt.Println(counter.Remove("f"))

	fmt.Println(counter.ToSlice())
	fmt.Println(counter.Length())

	counter.Reset()
	fmt.Println(counter.ToSlice())
	fmt.Println(counter.Length())

	// Output:
	// 43658
	// 2
	// 1632
	// 1
	// 1590
	// 234
	// 7435
	// 4361
	// 57279
	// a 43658
	// b 1
	// map[a:43658] 43658
	// map[b:1] 1
	// [{b 1} {d 234} {c 1590} {f 4361} {e 7435} {a 43658}]
	// [{f 4361} {e 7435} {a 43658}]
	// [{b 1} {d 234} {c 1590}]
	// 6
	// 4361
	// [{b 1} {d 234} {c 1590} {e 7435} {a 43658}]
	// 5
	// []
	// 0
}
