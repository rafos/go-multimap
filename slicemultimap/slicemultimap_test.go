package slicemultimap

import (
	"fmt"
	"github.com/rafos/go-multimap"

	"testing"
)

func TestClear(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a")

	if actualValue := m.Size(); actualValue != 8 {
		t.Errorf("expected %v, got %v", 8, actualValue)
	}
	if actualEmpty := m.Empty(); actualEmpty != false {
		t.Errorf("expected an empty multimap: %v, got %v", false, actualEmpty)
	}

	m.Clear()

	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("expected %v, got %v", 0, actualValue)
	}
	if actualEmpty := m.Empty(); actualEmpty != true {
		t.Errorf("expected an empty multimap: %v, got %v", true, actualEmpty)
	}
}
func TestPut(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a")

	if actualValue := m.Size(); actualValue != 8 {
		t.Errorf("expected %v, got %v", 8, actualValue)
	}
	if actualValue, expectedValue := m.Keys(), []int{1, 1, 2, 3, 4, 5, 6, 7}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.KeySet(), []int{1, 2, 3, 4, 5, 6, 7}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.Values(), []string{"a", "b", "c", "d", "e", "f", "g", "x"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}

	var expectedValue = []multimap.Entry[int, string]{
		{Key: 1, Value: "a"},
		{Key: 1, Value: "x"},
		{Key: 2, Value: "b"},
		{Key: 3, Value: "c"},
		{Key: 4, Value: "d"},
		{Key: 5, Value: "e"},
		{Key: 6, Value: "f"},
		{Key: 7, Value: "g"},
	}
	if actualValue := m.Entries(); !sameEntries(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}

	tests := []struct {
		key           int
		expectedValue []string
		expectedFound bool
	}{
		{1, []string{"a", "x"}, true},
		{2, []string{"b"}, true},
		{3, []string{"c"}, true},
		{4, []string{"d"}, true},
		{5, []string{"e"}, true},
		{6, []string{"f"}, true},
		{7, []string{"g"}, true},
		{8, nil, false},
		{9, nil, false},
	}

	for i, test := range tests {
		actualValue, actualFound := m.Get(test.key)
		if !sameElements(actualValue, test.expectedValue) || actualFound != test.expectedFound {
			t.Errorf("test %d: expected %v, got: %v ", i+1, test.expectedValue, actualValue)
		}
	}
}

func TestPutAll(t *testing.T) {
	m := New[int, string]()
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(2, "b")
	m.PutAll(1, []string{"a", "x", "y"})

	if actualValue := m.Size(); actualValue != 6 {
		t.Errorf("expected %v, got %v", 6, actualValue)
	}
	if actualValue, expectedValue := m.Keys(), []int{1, 1, 1, 2, 3, 4}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.KeySet(), []int{1, 2, 3, 4}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.Values(), []string{"a", "b", "c", "d", "x", "y"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}

	var expectedValue = []multimap.Entry[int, string]{
		{Key: 1, Value: "a"},
		{Key: 1, Value: "x"},
		{Key: 1, Value: "y"},
		{Key: 2, Value: "b"},
		{Key: 3, Value: "c"},
		{Key: 4, Value: "d"},
	}
	if actualValue := m.Entries(); !sameEntries(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}

	tests := []struct {
		key           int
		expectedValue []string
		expectedFound bool
	}{
		{1, []string{"a", "x", "y"}, true},
		{2, []string{"b"}, true},
		{3, []string{"c"}, true},
		{4, []string{"d"}, true},
		{5, nil, false},
		{6, nil, false},
	}

	for i, test := range tests {
		// Test for retrievals.
		actualValue, actualFound := m.Get(test.key)
		if !sameElements(actualValue, test.expectedValue) || actualFound != test.expectedFound {
			t.Errorf("test %d: expected %v, got: %v ", i+1, test.expectedValue, actualValue)
		}
	}
}

func TestContains(t *testing.T) {
	m := New[int, string]()
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(2, "b")
	m.PutAll(1, []string{"a", "x", "y"})

	if actualValue, expectedValue := m.Contains(1, "a"), true; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.Contains(1, "x"), true; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.Contains(1, "z"), false; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.ContainsKey(1), true; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.ContainsKey(5), false; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.ContainsValue("x"), true; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.ContainsValue("z"), false; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
}
func TestRemove(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a")

	m.Remove(5, "n")
	m.Remove(6, "f")
	m.Remove(7, "g")
	m.Remove(8, "h")
	m.Remove(5, "e")

	if actualValue := m.Size(); actualValue != 5 {
		t.Errorf("expected %v, got %v", 5, actualValue)
	}
	if actualValue, expectedValue := m.Keys(), []int{1, 1, 2, 3, 4}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.KeySet(), []int{1, 2, 3, 4}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := m.Values(), []string{"a", "b", "c", "d", "x"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}

	var expectedValue = []multimap.Entry[int, string]{
		{Key: 1, Value: "a"},
		{Key: 1, Value: "x"},
		{Key: 2, Value: "b"},
		{Key: 3, Value: "c"},
		{Key: 4, Value: "d"},
	}
	if actualValue := m.Entries(); !sameEntries(actualValue, expectedValue) {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}

	tests := []struct {
		key           int
		expectedValue []string
		expectedFound bool
	}{
		{1, []string{"a", "x"}, true},
		{2, []string{"b"}, true},
		{3, []string{"c"}, true},
		{4, []string{"d"}, true},
		{5, nil, false},
		{6, nil, false},
		{7, nil, false},
		{8, nil, false},
		{9, nil, false},
	}

	for i, test := range tests {
		actualValue, actualFound := m.Get(test.key)
		if !sameElements(actualValue, test.expectedValue) || actualFound != test.expectedFound {
			t.Errorf("test %d: expected %v, got: %v ", i+1, test.expectedValue, actualValue)
		}
	}

	m.Remove(1, "a")
	m.Remove(4, "d")
	m.Remove(1, "x")
	m.Remove(3, "c")
	m.Remove(2, "x")
	m.Remove(2, "b")

	if actualValue, expectedValue := fmt.Sprintf("%v", m.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%v", m.KeySet()), "[]"; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%v", m.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%v", m.Entries()), "[]"; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("expected %v, got %v", 0, actualValue)
	}
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("expected %v, got %v", true, actualValue)
	}
}

func TestRemoveAll(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a")

	m.RemoveAll(5)
	m.RemoveAll(6)
	m.RemoveAll(7)
	m.RemoveAll(8)
	m.RemoveAll(5)
	m.RemoveAll(1)
	m.RemoveAll(3)
	m.RemoveAll(2)
	m.RemoveAll(2)
	m.RemoveAll(4)
	m.RemoveAll(9)

	if actualValue, expectedValue := fmt.Sprintf("%v", m.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%v", m.KeySet()), "[]"; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%v", m.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%v", m.Entries()), "[]"; actualValue != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, actualValue)
	}
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("expected %v, got %v", 0, actualValue)
	}
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("expected %v, got %v", true, actualValue)
	}

	tests := []struct {
		key           int
		expectedValue []string
		expectedFound bool
	}{
		{1, nil, false},
		{2, nil, false},
		{3, nil, false},
		{4, nil, false},
		{5, nil, false},
		{6, nil, false},
		{7, nil, false},
		{8, nil, false},
		{9, nil, false},
	}

	for i, test := range tests {
		actualValue, actualFound := m.Get(test.key)
		if !sameElements(actualValue, test.expectedValue) || actualFound != test.expectedFound {
			t.Errorf("test %d: expected %v, got: %v ", i+1, test.expectedValue, actualValue)
		}
	}
}

// Helper function to check equality of keys/values.
func sameElements[V comparable](a []V, b []V) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Helper function to check equality of entries.
func sameEntries[K comparable, V comparable](a []multimap.Entry[K, V], b []multimap.Entry[K, V]) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av.Key == bv.Key && av.Value == bv.Value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Utilities for Benchmarking
func benchmarkGet(b *testing.B, m *MultiMap[any, any], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *MultiMap[any, any], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkPutAll(b *testing.B, m *MultiMap[any, any], size int) {
	v := make([]interface{}, 0)
	v = append(v, struct{}{})
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.PutAll(n, v)
		}
	}
}

func benchmarkRemove(b *testing.B, m *MultiMap[any, any], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n, struct{}{})
		}
	}
}

func benchmarkRemoveAll(b *testing.B, m *MultiMap[any, any], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.RemoveAll(n)
		}
	}
}

func BenchmarkMultiMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkMultiMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkMultiMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkMultiMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkMultiMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[any, any]()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkMultiMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkMultiMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkMultiMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkMultiMapPutAll100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[any, any]()
	b.StartTimer()
	benchmarkPutAll(b, m, size)
}

func BenchmarkMultiMapPutAll1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPutAll(b, m, size)
}

func BenchmarkMultiMapPutAll10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPutAll(b, m, size)
}

func BenchmarkMultiMapPutAll100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPutAll(b, m, size)
}

func BenchmarkMultiMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkMultiMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkMultiMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkMultiMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkMultiMapRemoveAll100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemoveAll(b, m, size)
}

func BenchmarkMultiMapRemoveAll1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemoveAll(b, m, size)
}

func BenchmarkMultiMapRemoveAll10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemoveAll(b, m, size)
}

func BenchmarkMultiMapRemoveAll100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[any, any]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemoveAll(b, m, size)
}
