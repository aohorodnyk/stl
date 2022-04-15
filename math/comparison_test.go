package math_test

import (
	"fmt"
	"math"
	"testing"

	mathstl "github.com/aohorodnyk/stl/math"
)

func TestMinByte(t *testing.T) {
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

	for idx, p := range prov {
		p := p

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Min(p.input...)
			if res != p.want {
				t.Errorf("got %v, expected %v", res, p.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
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

	for idx, p := range prov {
		p := p

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Min(p.input...)
			if res != p.want {
				t.Errorf("got %v, expected %v", res, p.want)
			}
		})
	}
}

func TestMinStr(t *testing.T) {
	t.Parallel()

	type str string

	prov := []struct {
		input []str
		want  str
	}{
		{
			input: []str{},
			want:  "",
		},
		{
			input: []str{"a"},
			want:  "a",
		},
		{
			input: []str{"", "1", "a", "b", "Z"},
			want:  "",
		},
		{
			input: []str{"4", "1", "a", "b", "Z"},
			want:  "1",
		},
		{
			input: []str{"A", "m", "a", "b", "Z"},
			want:  "A",
		},
		{
			input: []str{"cab", "caa", "cba"},
			want:  "caa",
		},
		{
			input: []str{"aaa", "aa2", "aa1", "a2a", "a1a", "a0a", "ava"},
			want:  "a0a",
		},
	}

	for idx, p := range prov {
		p := p

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Min(p.input...)
			if res != p.want {
				t.Errorf("got %v, expected %v", res, p.want)
			}
		})
	}
}

func TestMaxByte(t *testing.T) {
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

	for idx, p := range prov {
		p := p

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Max(p.input...)
			if res != p.want {
				t.Errorf("got %v, expected %v", res, p.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
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

	for idx, p := range prov {
		p := p

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Max(p.input...)
			if res != p.want {
				t.Errorf("got %v, expected %v", res, p.want)
			}
		})
	}
}

func TestMaxStr(t *testing.T) {
	t.Parallel()

	type str string

	prov := []struct {
		input []str
		want  str
	}{
		{
			input: []str{},
			want:  "",
		},
		{
			input: []str{"a"},
			want:  "a",
		},
		{
			input: []str{"", "1", "a", "b", "Z"},
			want:  "b",
		},
		{
			input: []str{"4", "1", "a", "b", "Z"},
			want:  "b",
		},
		{
			input: []str{"A", "m", "a", "b", "Z"},
			want:  "m",
		},
		{
			input: []str{"cab", "caa", "cba"},
			want:  "cba",
		},
		{
			input: []str{"aaa", "aa2", "aa1", "a2a", "a1a", "a0a", "ava"},
			want:  "ava",
		},
	}

	for idx, p := range prov {
		p := p

		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			t.Parallel()

			res := mathstl.Max(p.input...)
			if res != p.want {
				t.Errorf("got %v, expected %v", res, p.want)
			}
		})
	}
}
