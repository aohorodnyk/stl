package collections_test

import (
	"fmt"
	"testing"

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

func TestInc(t *testing.T) {
	t.Parallel()

	counter := collections.NewCounter[uint]()

	counter.Inc(1)
	counter.Inc(1)
	counter.Inc(1)
	counter.Inc(6)
	counter.Inc(0)
	counter.Inc(12534)
	counter.Inc(12534)
	counter.Inc(47652)
	counter.Inc(47652)
	counter.Inc(47652)
	counter.Inc(47652)
	counter.Inc(47652)
	counter.Inc(47652)

	if counter.Counted(1) != 3 {
		t.Errorf("Expected 3, got %d", counter.Counted(1))
	}

	if counter.Counted(6) != 1 {
		t.Errorf("Expected 1, got %d", counter.Counted(6))
	}

	if counter.Counted(0) != 1 {
		t.Errorf("Expected 1, got %d", counter.Counted(0))
	}

	if counter.Counted(12534) != 2 {
		t.Errorf("Expected 2, got %d", counter.Counted(12534))
	}

	if counter.Counted(47652) != 6 {
		t.Errorf("Expected 6, got %d", counter.Counted(47652))
	}

	if counter.Counted(47651) != 0 {
		t.Errorf("Expected 0, got %d", counter.Counted(47652))
	}

	if counter.Counted(47653) != 0 {
		t.Errorf("Expected 0, got %d", counter.Counted(47652))
	}

	if counter.Sum() != 13 {
		t.Errorf("Expected 23, got %d", counter.Sum())
	}
}

func TestDec(t *testing.T) {
	t.Parallel()

	counter := collections.NewCounter[int]()

	counter.Dec(1)
	counter.Dec(2)

	counter.Add(-5, 5)
	counter.Dec(-5)

	if counter.Counted(1) != 0 {
		t.Errorf("Expected 0, got %d", counter.Counted(1))
	}

	if counter.Counted(2) != 0 {
		t.Errorf("Expected 0, got %d", counter.Counted(2))
	}

	if counter.Counted(-5) != 4 {
		t.Errorf("Expected 4, got %d", counter.Counted(-5))
	}
}

func TestAddSub(t *testing.T) {
	t.Parallel()

	counter := collections.NewCounter[bool]()

	counter.Add(true, 23535234)
	counter.Add(true, 1632)
	counter.Sub(true, 1000)

	if counter.Counted(true) != 23535866 {
		t.Errorf("Expected 23535866, got %d", counter.Counted(true))
	}

	counter.Add(false, 25)
	counter.Add(false, 18)
	counter.Add(false, 7)
	counter.Sub(false, 5)

	if counter.Counted(false) != 45 {
		t.Errorf("Expected 45, got %d", counter.Counted(false))
	}

	if counter.Sum() != 23535911 {
		t.Errorf("Expected 23535911, got %d", counter.Sum())
	}

	counter.Sub(true, 23535865)

	if counter.Counted(true) != 1 {
		t.Errorf("Expected 1, got %d", counter.Counted(true))
	}

	counter.Sub(true, 1)

	if counter.Counted(true) != 0 {
		t.Errorf("Expected 0, got %d", counter.Counted(true))
	}

	counter.Sub(false, 75)

	if counter.Counted(false) != 0 {
		t.Errorf("Expected 0, got %d", counter.Counted(false))
	}

	if len(counter.ToSlice()) != 0 {
		t.Errorf("Expected 0, got %d", len(counter.ToSlice()))
	}
}
