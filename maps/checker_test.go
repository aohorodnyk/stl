package maps_test

import (
	"testing"

	"github.com/aohorodnyk/stl/maps"
	"github.com/aohorodnyk/stl/types"
)

func TestKeyExists(t *testing.T) {
	t.Parallel()

	if maps.KeyExists[string, types.Empty](nil, "unknown key") {
		t.Error("KeyExists returned true, but 'unknown key' cannot be found in nil map")
	}

	if !maps.KeyExists(map[int]types.Empty{4: {}, 2: {}}, 2) {
		t.Error("KeyExists returned false, but 2 was set as a key")
	}

	if maps.KeyExists(map[int]types.Empty{4: {}, 2: {}}, 3) {
		t.Error("KeyExists returned true, but 3 was not set as a key")
	}

	if !maps.KeyExists(map[string]int{"known key": 2, "key2": 1}, "known key") {
		t.Error("KeyExists returned false, but 'known key' was set as a key")
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()

	containerInt := map[string]int{"one": 1, "two": 2, "three": 3}
	searchInt := 2

	if key, ok := maps.Search(containerInt, searchInt); !ok || key != "two" {
		t.Errorf("Expected value to be found, but got ok: %v, key: %s", ok, key)
	}

	if key, ok := maps.Search(containerInt, -searchInt); ok || key != "" {
		t.Errorf("Expected value to not be found, but got ok: %v, key: %s", ok, key)
	}

	containerStr := map[int]string{10: "ten", 20: "twenty", 100: "one hundred"}
	searchStr := "one hundred"

	if key, ok := maps.Search(containerStr, searchStr); !ok || key != 100 {
		t.Errorf("Expected value to be found, but got ok: %v, key: %d", ok, key)
	}

	if key, ok := maps.Search(containerStr, searchStr+"_test"); ok || key != 0 {
		t.Errorf("Expected value to not be found, but got ok: %v, key: %d", ok, key)
	}
}

func TestSearchFun(t *testing.T) {
	t.Parallel()

	container := map[string]int{"one": 1, "two": 2, "three": 3}
	searchVal := 2
	cmpFound := func(key string) bool {
		return container[key] == searchVal
	}
	cmpNotFound := func(key string) bool {
		return false
	}

	if key, ok := maps.SearchFunc(container, cmpFound); !ok || key != "two" {
		t.Errorf("Expected value to be found, but got ok: %v, key: %s", ok, key)
	}

	if key, ok := maps.SearchFunc(container, cmpNotFound); ok || key != "" {
		t.Errorf("Expected value to be found, but got ok: %v, key: %s", ok, key)
	}
}
