package math_test

import (
	"fmt"
	"math"
	"testing"

	mathstl "github.com/aohorodnyk/stl/math"
)

func TestMinMultiByte(t *testing.T) {
	t.Parallel()

	prov := []struct {
		input []byte
		want  byte
	}{
		{
			input: []byte{},
			want:  0,
		},
		{
			input: []byte{1},
			want:  1,
		},
		{
			input: []byte{5, 0, math.MaxUint8, 234},
			want:  0,
		},
		{
			input: []byte{5, 3, 2, 123, 245, 255, 1, 12, 23},
			want:  1,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, provData := range prov {
		provData := provData

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.MinMulti(provData.input...)
			if res != provData.want {
				t.Errorf("got %v, expected %v", res, provData.want)
			}
		})
	}
}

func TestMinMultiInt(t *testing.T) {
	t.Parallel()

	type myInt int

	prov := []struct {
		input []myInt
		want  myInt
	}{
		{
			input: []myInt{},
			want:  0,
		},
		{
			input: []myInt{1},
			want:  1,
		},
		{
			input: []myInt{5, 0, math.MaxUint8, 234},
			want:  0,
		},
		{
			input: []myInt{5, 3, 2, 123, 245, math.MinInt, math.MaxInt, 1, 12, 23},
			want:  math.MinInt,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, provData := range prov {
		provData := provData

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.MinMulti(provData.input...)
			if res != provData.want {
				t.Errorf("got %v, expected %v", res, provData.want)
			}
		})
	}
}

func TestMaxMultiByte(t *testing.T) {
	t.Parallel()

	prov := []struct {
		input []byte
		want  byte
	}{
		{
			input: []byte{},
			want:  0,
		},
		{
			input: []byte{1},
			want:  1,
		},
		{
			input: []byte{5, 0, math.MaxUint8, 234},
			want:  math.MaxUint8,
		},
		{
			input: []byte{5, 3, 2, 123, 245, 1, 12, 23},
			want:  245,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, provData := range prov {
		provData := provData

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.MaxMulti(provData.input...)
			if res != provData.want {
				t.Errorf("got %v, expected %v", res, provData.want)
			}
		})
	}
}

func TestMaxMultiInt(t *testing.T) {
	t.Parallel()

	type myInt int

	prov := []struct {
		input []myInt
		want  myInt
	}{
		{
			input: []myInt{},
			want:  0,
		},
		{
			input: []myInt{1},
			want:  1,
		},
		{
			input: []myInt{5, 0, math.MaxUint8, 234},
			want:  math.MaxUint8,
		},
		{
			input: []myInt{5, 3, 2, 123, 245, math.MinInt, math.MaxInt, 1, 12, 23},
			want:  math.MaxInt,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, provData := range prov {
		provData := provData

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.MaxMulti(provData.input...)
			if res != provData.want {
				t.Errorf("got %v, expected %v", res, provData.want)
			}
		})
	}
}

func TestMinByte(t *testing.T) {
	t.Parallel()

	prov := []struct {
		a, b byte
		want byte
	}{
		{
			a:    0,
			b:    0,
			want: 0,
		},
		{
			a:    1,
			b:    1,
			want: 1,
		},
		{
			a:    255,
			b:    0,
			want: 0,
		},
		{
			a:    1,
			b:    234,
			want: 1,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, provData := range prov {
		provData := provData

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Min(provData.a, provData.b)
			if res != provData.want {
				t.Errorf("got %v, expected %v", res, provData.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
	t.Parallel()

	type myInt int

	prov := []struct {
		a, b myInt
		want myInt
	}{
		{
			a:    0,
			b:    0,
			want: 0,
		},
		{
			a:    1,
			b:    1,
			want: 1,
		},
		{
			a:    235214,
			b:    234,
			want: 234,
		},
		{
			a:    math.MinInt,
			b:    877423564,
			want: math.MinInt,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, provData := range prov {
		provData := provData

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Min(provData.a, provData.b)
			if res != provData.want {
				t.Errorf("got %v, expected %v", res, provData.want)
			}
		})
	}
}

func TestMaxByte(t *testing.T) {
	t.Parallel()

	prov := []struct {
		a, b byte
		want byte
	}{
		{
			a:    0,
			b:    0,
			want: 0,
		},
		{
			a:    1,
			b:    1,
			want: 1,
		},
		{
			a:    231,
			b:    math.MaxUint8,
			want: math.MaxUint8,
		},
		{
			a:    245,
			b:    123,
			want: 245,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, provData := range prov {
		provData := provData

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Max(provData.a, provData.b)
			if res != provData.want {
				t.Errorf("got %v, expected %v", res, provData.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	t.Parallel()

	type myInt int

	prov := []struct {
		a, b myInt
		want myInt
	}{
		{
			a:    0,
			b:    0,
			want: 0,
		},
		{
			a:    1,
			b:    1,
			want: 1,
		},
		{
			a:    math.MaxInt,
			b:    236523452,
			want: math.MaxInt,
		},
		{
			a:    123,
			b:    math.MaxInt,
			want: math.MaxInt,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, provData := range prov {
		provData := provData

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Max(provData.a, provData.b)
			if res != provData.want {
				t.Errorf("got %v, expected %v", res, provData.want)
			}
		})
	}
}

func TestCompareMultiCustom(t *testing.T) {
	t.Parallel()

	type customType struct {
		field int
	}

	cmpMin := func(first, second customType) bool { return first.field < second.field }
	cmpMax := func(first, second customType) bool { return first.field > second.field }
	cmpRef := func(first, second *customType) bool { return first.field > second.field }

	min := mathstl.CompareMulti(cmpMin, customType{field: 1}, customType{field: 2}, customType{field: -15}, customType{field: -2})
	if min.field != -15 {
		t.Errorf("got %v, expected %v", min.field, -15)
	}

	max := mathstl.CompareMulti(cmpMax, customType{field: 1}, customType{field: 2}, customType{field: -15}, customType{field: -2})
	if max.field != 2 {
		t.Errorf("got %v, expected %v", max.field, 2)
	}

	empty := mathstl.CompareMulti(cmpMax)
	if empty.field != 0 {
		t.Errorf("got %v, expected %v", max.field, 0)
	}

	emptyRef := mathstl.CompareMulti(cmpRef)
	if emptyRef != nil {
		t.Errorf("got %v, expected %v", emptyRef, nil)
	}

	maxRef := mathstl.CompareMulti(cmpRef, &customType{field: 1}, &customType{field: 2}, &customType{field: -15}, &customType{field: -2})
	if maxRef.field != 2 {
		t.Errorf("got %v, expected %v", emptyRef, 2)
	}

	maxRef = mathstl.CompareMulti(cmpRef, &customType{field: 1})
	if maxRef.field != 1 {
		t.Errorf("got %v, expected %v", emptyRef, 1)
	}
}
