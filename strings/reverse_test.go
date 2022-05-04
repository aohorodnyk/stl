package strings_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/strings"
)

func ExampleReverse() {
	fmt.Println(strings.Reverse(""))
	fmt.Println(strings.Reverse("w"))
	fmt.Println(strings.Reverse("wa"))
	fmt.Println(strings.Reverse("was"))
	fmt.Println(strings.Reverse("eeSee"))
	fmt.Println(strings.Reverse("Hello, world"))

	// Output:
	//
	// w
	// aw
	// saw
	// eeSee
	// dlrow ,olleH
}
