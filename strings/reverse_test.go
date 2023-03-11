package strings_test

import (
	"fmt"

	"github.com/aohorodnyk/stl/strings"
)

func ExampleReverse() {
	empty := strings.Reverse("")
	fmt.Printf("%T: '%s'\n", empty, empty)

	oneLetter := strings.Reverse("a")
	fmt.Printf("%T: '%s'\n", oneLetter, oneLetter)

	twoLetters := strings.Reverse("ab")
	fmt.Printf("%T: '%s'\n", twoLetters, twoLetters)

	threeLetters := strings.Reverse("abc")
	fmt.Printf("%T: '%s'\n", threeLetters, threeLetters)

	type inheritedType string

	stringCustom := strings.Reverse(inheritedType("custom"))
	fmt.Printf("%T: '%s'\n", stringCustom, stringCustom)

	// Output:
	// string: ''
	// string: 'a'
	// string: 'ba'
	// string: 'cba'
	// strings_test.inheritedType: 'motsuc'
}
