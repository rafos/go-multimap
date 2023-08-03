package multimap

// Entry represents a key/value pair inside a multimap.
type Entry[K comparable, V comparable] struct {
	Key   K
	Value V
}

// MultiMap interface that all multimaps implement.
type MultiMap[K comparable, V comparable] interface {
	Get(key K) (value []V, found bool)

	Put(key K, value V)
	PutAll(key K, value []V)

	Remove(key K, value V)
	RemoveAll(key K)

	Contains(key K, value V) bool
	ContainsKey(key K) bool
	ContainsValue(value V) bool

	Entries() []Entry[K, V]
	Keys() []K
	KeySet() []K
	Values() []V

	Clear()
	Empty() bool
	Size() int
}
