package main

import "testing"

func TestIsTrue(t *testing.T) {
	if !IsTrue() {
		t.Fatal("Something Failed")
	}
}
