package bin_test

import (
	"fmt"
	"math"

	"github.com/aohorodnyk/stl/math/bin"
)

func ExampleIsSetBit() {
	flagInt8 := int8(-1)
	fmt.Println(bin.IsSetBit(flagInt8, 0))
	fmt.Println(bin.IsSetBit(flagInt8, 7))
	fmt.Println(bin.IsSetBit(flagInt8, 8))
	fmt.Println(bin.IsSetBit(flagInt8, 9))

	flagInt8 = 0
	fmt.Println(bin.IsSetBit(flagInt8, 0))

	flagInt8 = 1
	fmt.Println(bin.IsSetBit(flagInt8, 0))

	flagInt8 = 1 << 2
	fmt.Println(bin.IsSetBit(flagInt8, 1))
	fmt.Println(bin.IsSetBit(flagInt8, 2))
	fmt.Println(bin.IsSetBit(flagInt8, 3))

	flagUint32 := uint32(math.MaxUint32)
	fmt.Println(bin.IsSetBit(flagUint32, 0))
	fmt.Println(bin.IsSetBit(flagUint32, 31))
	fmt.Println(bin.IsSetBit(flagUint32, 32))

	flagInt64 := int64(-1)
	fmt.Println(bin.IsSetBit(flagInt64, 0))
	fmt.Println(bin.IsSetBit(flagInt64, 63))
	fmt.Println(bin.IsSetBit(flagInt64, 64))

	// Output:
	// true
	// true
	// false
	// false
	// false
	// true
	// false
	// true
	// false
	// true
	// true
	// false
	// true
	// true
	// false
}

func ExampleSetBit() {
	int8s := bin.SetBit(int8(0), 0)
	fmt.Println(int8s)
	int8s = bin.SetBit(int8s, 1)
	fmt.Println(int8s)
	int8s = bin.SetBit(int8s, 1)
	fmt.Println(int8s)
	int8s = bin.SetBit(int8s, 6)
	fmt.Println(int8s)
	int8s = bin.SetBit(int8s, 7)
	fmt.Println(int8s)
	int8s = bin.SetBit(int8s, 8)
	fmt.Println(int8s)

	int8u := bin.SetBit(uint8(0), 0)
	fmt.Println(int8u)
	int8u = bin.SetBit(int8u, 1)
	fmt.Println(int8u)
	int8u = bin.SetBit(int8u, 1)
	fmt.Println(int8u)
	int8u = bin.SetBit(int8u, 6)
	fmt.Println(int8u)
	int8u = bin.SetBit(int8u, 7)
	fmt.Println(int8u)
	int8u = bin.SetBit(int8u, 8)
	fmt.Println(int8u)

	int64s := bin.SetBit(int64(0), 0)
	fmt.Println(int64s)
	int64s = bin.SetBit(int64s, 1)
	fmt.Println(int64s)
	int64s = bin.SetBit(int64s, 1)
	fmt.Println(int64s)
	int64s = bin.SetBit(int64s, 6)
	fmt.Println(int64s)
	int64s = bin.SetBit(int64s, 60)
	fmt.Println(int64s)
	int64s = bin.SetBit(int64s, 61)
	fmt.Println(int64s)
	int64s = bin.SetBit(int64s, 62)
	fmt.Println(int64s)
	int64s = bin.SetBit(int64s, 63)
	fmt.Println(int64s)
	int64s = bin.SetBit(int64s, 64)
	fmt.Println(int64s)

	// Output:
	// 1
	// 3
	// 3
	// 67
	// -61
	// -61
	// 1
	// 3
	// 3
	// 67
	// 195
	// 195
	// 1
	// 3
	// 3
	// 67
	// 1152921504606847043
	// 3458764513820540995
	// 8070450532247928899
	// -1152921504606846909
	// -1152921504606846909
}

func ExampleUnsetBit() {
	int8s := bin.UnsetBit(int8(0), 0)
	fmt.Println(int8s)
	int8s = bin.UnsetBit(int8(1), 0)
	fmt.Println(int8s)
	int8s = bin.UnsetBit(int8(7), 1)
	fmt.Println(int8s)
	int8s = bin.UnsetBit(int8(5), 1)
	fmt.Println(int8s)
	int8s = bin.UnsetBit(int8(-1), 6)
	fmt.Println(int8s)
	int8s = bin.UnsetBit(int8(-1), 7)
	fmt.Println(int8s)

	int64u := bin.UnsetBit(uint64(0), 0)
	fmt.Println(int64u)
	int64u = bin.UnsetBit(uint64(267435374365236435), 53)
	fmt.Println(int64u)
	int64u = bin.UnsetBit(uint64(math.MaxUint64), 62)
	fmt.Println(int64u)
	int64u = bin.UnsetBit(uint64(math.MaxUint64), 60)
	fmt.Println(int64u)

	// Output:
	// 0
	// 0
	// 5
	// 5
	// -65
	// 127
	// 0
	// 258428175110495443
	// 13835058055282163711
	// 17293822569102704639
}
