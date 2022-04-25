package reflect_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/stl/reflect"
)

func TestDeepEqual(t *testing.T) {
	t.Parallel()

	provider := []struct {
		a, b interface{}
		exp  bool
	}{
		{
			a:   nil,
			b:   nil,
			exp: true,
		},
		{
			a:   nil,
			b:   1,
			exp: false,
		},
		{
			a:   1,
			b:   nil,
			exp: false,
		},
		{
			a:   1,
			b:   1,
			exp: true,
		},
		{
			a:   "asdfewf",
			b:   "gwdfcv",
			exp: false,
		},
		{
			a:   "asdfewf",
			b:   "asdfewf",
			exp: true,
		},
		{
			a:   0,
			b:   0,
			exp: true,
		},
		{
			a:   124235.234523,
			b:   124235.234523,
			exp: true,
		},
		{
			a:   124235.0,
			b:   124235.0,
			exp: true,
		},
		{
			a:   124235.0,
			b:   124235,
			exp: false,
		},
		{
			a:   []string{"asdf", "asdf", "asdf"},
			b:   []string{"asdf", "asdf", "asdf"},
			exp: true,
		},
		{
			a:   []string{"asdf", "asdf", "asda"},
			b:   []string{"asdf", "asdf", "asdf"},
			exp: false,
		},
		{
			a:   map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			b:   map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			exp: true,
		},
		{
			a:   map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwe2r"},
			b:   map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			exp: false,
		},
	}

	//nolint:paralleltest // We use provider data in subtests runs.
	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			t.Parallel()

			act := reflect.DeepEqual(prov.a, prov.b)

			if act != prov.exp {
				t.Errorf("expected %v, got %v", prov.exp, act)
			}
		})
	}
}

func TestDeepEqualCpm(t *testing.T) {
	t.Parallel()

	provider := []struct {
		a, b       interface{}
		comparable bool
		exp        bool
	}{
		{
			a:          nil,
			b:          1,
			comparable: false,
			exp:        false,
		},
		{
			a:          1,
			b:          nil,
			comparable: false,
			exp:        false,
		},
		{
			a:          nil,
			b:          1,
			comparable: true,
			exp:        false,
		},
		{
			a:          1,
			b:          nil,
			comparable: true,
			exp:        false,
		},
		{
			a:          nil,
			b:          nil,
			comparable: false,
			exp:        true,
		},
		{
			a:          nil,
			b:          nil,
			comparable: true,
			exp:        true,
		},
		{
			a:          1,
			b:          1,
			comparable: true,
			exp:        true,
		},
		{
			a:          "asdfewf",
			b:          "gwdfcv",
			comparable: true,
			exp:        false,
		},
		{
			a:          "asdfewf",
			b:          "asdfewf",
			comparable: true,
			exp:        true,
		},
		{
			a:          0,
			b:          0,
			comparable: true,
			exp:        true,
		},
		{
			a:          124235.234523,
			b:          124235.234523,
			comparable: true,
			exp:        true,
		},
		{
			a:          124235.0,
			b:          124235.0,
			comparable: true,
			exp:        true,
		},
		{
			a:          124235.0,
			b:          124235,
			comparable: true,
			exp:        false,
		},
		{
			a:          124235.0,
			b:          124235,
			comparable: false,
			exp:        false,
		},
		{
			a:          124235.0,
			b:          "124235",
			comparable: false,
			exp:        false,
		},
		{
			a:          []string{"asdf", "asdf", "asdf"},
			b:          []string{"asdf", "asdf", "asdf"},
			comparable: false,
			exp:        true,
		},
		{
			a:          []string{"asdf", "asdf", "asda"},
			b:          []string{"asdf", "asdf", "asdf"},
			comparable: false,
			exp:        false,
		},
		{
			a:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			b:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			comparable: false,
			exp:        true,
		},
		{
			a:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwe2r"},
			b:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			comparable: false,
			exp:        false,
		},
		{
			a:          []string{"asdf", "qwetr", "berwf", "qwet21", "wfwe", "efqwe2r"},
			b:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			comparable: false,
			exp:        false,
		},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			t.Parallel()

			act := reflect.DeepEqualCmp(prov.comparable, prov.a, prov.b)

			if act != prov.exp {
				t.Errorf("expected %v, got %v", prov.exp, act)
			}
		})
	}
}

func TestDeepEqualCpmPanic(t *testing.T) {
	t.Parallel()

	provider := []struct {
		a, b       interface{}
		comparable bool
		exp        bool
		panic      bool
	}{
		{
			a:          1,
			b:          1,
			comparable: true,
			exp:        true,
			panic:      false,
		},
		{
			a:          "asdfewf",
			b:          "gwdfcv",
			comparable: true,
			exp:        false,
			panic:      false,
		},
		{
			a:          "asdfewf",
			b:          "asdfewf",
			comparable: true,
			exp:        true,
			panic:      false,
		},
		{
			a:          0,
			b:          0,
			comparable: true,
			exp:        true,
			panic:      false,
		},
		{
			a:          34524,
			b:          "34524",
			comparable: true,
			exp:        false,
			panic:      false,
		},
		{
			a:          124235,
			b:          124235.234523,
			comparable: true,
			exp:        false,
			panic:      false,
		},
		{
			a:          124235.234523,
			b:          124235.234523,
			comparable: true,
			exp:        true,
			panic:      false,
		},
		{
			a:          124235.0,
			b:          124235.0,
			comparable: true,
			exp:        true,
			panic:      false,
		},
		{
			a:          124235.0,
			b:          124235,
			comparable: true,
			exp:        false,
			panic:      false,
		},
		{
			a:          []string{"asdf", "asdf", "asdf"},
			b:          []string{"asdf", "asdf", "asdf"},
			comparable: false,
			exp:        true,
			panic:      false,
		},
		{
			a:          234,
			b:          map[string]string{"asdf": "qwetr"},
			comparable: true,
			exp:        false,
			panic:      false,
		},
		{
			a:          []string{"asdf", "asdf", "asda"},
			b:          map[string]string{"asdf": "qwetr"},
			comparable: true,
			exp:        false,
			panic:      false,
		},
		{
			a:          []string{"asdf", "asdf", "asda"},
			b:          []string{"asdf", "asdf", "asdf"},
			comparable: true,
			exp:        false,
			panic:      true,
		},
		{
			a:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			b:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			comparable: false,
			exp:        true,
			panic:      false,
		},
		{
			a:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwe2r"},
			b:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			comparable: false,
			exp:        false,
			panic:      false,
		},
		{
			a:          []string{"asdf", "asdf", "asdf"},
			b:          []string{"asdf", "asdf", "asdf"},
			comparable: true,
			exp:        true,
			panic:      true,
		},
		{
			a:          []string{"asdf", "asdf", "asda"},
			b:          []string{"asdf", "asdf", "asdf"},
			comparable: true,
			exp:        false,
			panic:      true,
		},
		{
			a:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			b:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			comparable: true,
			exp:        true,
			panic:      true,
		},
		{
			a:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwe2r"},
			b:          map[string]string{"asdf": "qwetr", "berwf": "qwet21", "wfwe": "efqwetr"},
			comparable: true,
			exp:        false,
			panic:      true,
		},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			t.Parallel()

			defer func() {
				err := recover()
				if (err != nil) != prov.panic {
					t.Errorf("expected panic: %v, got: %v", prov.panic, err)
				}
			}()

			act := reflect.DeepEqualCmp(prov.comparable, prov.a, prov.b)

			if act != prov.exp {
				t.Errorf("expected %v, got %v", prov.exp, act)
			}
		})
	}
}
