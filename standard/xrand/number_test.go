package xrand

import "testing"

func TestRand(t *testing.T) {
	for i := 1; i <= 5; i++ {
		t.Log(Int(1, 2))
	}
	for i := 1; i <= 5; i++ {
		t.Log(Float(1, 2, 2))
	}
	for i := 1; i <= 5; i++ {
		t.Log(String(5))
	}
}
