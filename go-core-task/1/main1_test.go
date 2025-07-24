package main

import "testing"

func TestCorrectHash_GenerateHashedString(t *testing.T) {
	data := []interface{}{1, 052, 0x52, 2.5, "hello", true, complex(1, 2)}
	salt := "go-2024"
	expectedHash := "a481f11da31065bb90824b8c3208a3962a1ddeea5c2e70a91d6b2adbbf8a725a"
	result := GenerateHashedString(data, salt)
	if result != expectedHash {
		t.Errorf("expected hash %s, but got %s", expectedHash, result)
	}
}

func TestNilData_GenerateHashedString(t *testing.T) {
	var data []interface{}
	salt := "go-2024"
	result := GenerateHashedString(data, salt)
	if result == "" {
		t.Error("expected non-empty hash for empty input, but got an empty string")
	}
}

func TestSameHash_GenerateHashedString(t *testing.T) {
	data := []interface{}{1, 052, 0x52, 2.5, "hello", true, complex(1, 2)}
	salt := "go-2024"
	result1 := GenerateHashedString(data, salt)
	result2 := GenerateHashedString(data, salt)
	if result1 != result2 {
		t.Errorf("expected the same hash for the same input, but got %s and %s", result1, result2)
	}
}
