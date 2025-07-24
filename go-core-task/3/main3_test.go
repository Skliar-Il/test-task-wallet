package main

import (
	"reflect"
	"testing"
)

func TestAddAndExist_StringIntMap(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 10)
	m.Add("b", 20)

	if !m.Exists("a") {
		t.Errorf("key 'a' should exist")
	}
	if !m.Exists("b") {
		t.Errorf("key 'b' should exist")
	}
	if m.Exists("c") {
		t.Errorf("key 'c' should not exist")
	}
}

func TestGet_StringIntMap(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 10)

	v, ok := m.Get("a")
	if !ok || v != 10 {
		t.Errorf("expected true 10, got %d, ok=%v", v, ok)
	}

	_, ok = m.Get("c")
	if ok {
		t.Errorf("expected false for non-existing key 'c'")
	}
}

func TestRemove_StringIntMap(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 10)
	m.Remove("a")

	if m.Exists("a") {
		t.Errorf("key 'a' should be removed")
	}
}

func TestCopy_StringIntMap(t *testing.T) {
	m := NewStringIntMap()
	m.Add("b", 20)

	cpy := m.Copy()
	expected := map[string]int{"b": 20}
	if !reflect.DeepEqual(cpy, expected) {
		t.Errorf("expected copy %v, got %v", expected, cpy)
	}

	m.Add("d", 30)
	if _, ok := cpy["d"]; ok {
		t.Errorf("copy should not contain new key 'd'")
	}
}
