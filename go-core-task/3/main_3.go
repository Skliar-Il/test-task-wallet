package main

import (
	"fmt"
)

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	cpy := make(map[string]int, len(m.data))
	for k, v := range m.data {
		cpy[k] = v
	}
	return cpy
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]
	return val, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("one", 1)
	fmt.Println(m.Exists("one"))
	v, ok := m.Get("one")
	fmt.Println(v, ok)
	m.Remove("one")
	fmt.Println(m.Exists("one"))
}
