package main

import "testing"

func TestABCheck(t *testing.T) {
	if !ABCheck("123adcvb") {
		t.Error("Expected true found false")
	}
	if !ABCheck("123adzvb") {
		t.Error("Expected true found false")
	}
	if !ABCheck("abccccazzzb") {
		t.Error("Expected true found false")
	}
	if !ABCheck("bzzza") {
		t.Error("Expected true found false")
	}
	if !ABCheck("lane borrowed") {
		t.Error("Expected true found false")
	}

	if ABCheck("lne borroweda") {
		t.Error("Expected false found true")
	}
}
