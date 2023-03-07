package types

// Empty alias to empty struct{} type.
// This type helps to avoid extra brackets {}.
// Example:
//
//	printSet := func(set map[string]Empty) {
//		fmt.Print(set)
//	}
//
//	a := struct{}{}
//	b := Empty{}
//	printSet(a)
//	printSet(b)
//
//	Outputs:
//	map[a:{} b:{}]
//	map[c:{} d:{}]
type Empty = struct{}
