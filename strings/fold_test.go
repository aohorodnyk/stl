package strings_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/stl/strings"
)

func TestEqualFold(t *testing.T) {
	t.Parallel()

	if !strings.EqualFold("Hello", "hello") {
		t.Errorf("Failed to compare strings with EqualFold")
	}
}

func TestHasPrefixFold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s      string
		prefix string
		result bool
	}{
		{"Hello", "hello", true},
		{"Hello", "Hello", true},
		{"Hello", "Hello world", false},
		{"Hello", "world", false},
		{"Hello", "", true},
		{"", "", true},
		{"", "Hello", false},
		{"Halloween day", "hallo", true},
	}

	for idx, test := range tests {
		test := test

		t.Run(fmt.Sprintf("TestHasPrefixFold_%d", idx), func(t *testing.T) {
			t.Parallel()

			if strings.HasPrefixFold(test.s, test.prefix) != test.result {
				t.Errorf("Failed to compare strings with HasPrefixFold")
			}
		})
	}
}

func TestHasSuffixFold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s      string
		suffix string
		result bool
	}{
		{"Hello", "hello", true},
		{"Hello", "Hello", true},
		{"Hello", "Hello world", false},
		{"Hello", "world", false},
		{"Hello", "", true},
		{"", "", true},
		{"", "Hello", false},
		{"Halloween day", "day", true},
	}

	for idx, test := range tests {
		test := test

		t.Run(fmt.Sprintf("TestHasSuffixFold_%d", idx), func(t *testing.T) {
			t.Parallel()

			if strings.HasSuffixFold(test.s, test.suffix) != test.result {
				t.Errorf("Failed to compare strings with HasSuffixFold")
			}
		})
	}
}
