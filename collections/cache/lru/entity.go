package lru

type entity[Key comparable, Value any] struct {
	key   Key
	value Value
}
