package stringset

import (
	"testing"
)

func TestStringset (t *testing.T) {
	ss := New()
	ss.Add("a")
	if !ss.Contains("a") {
		t.Fatal("ss does not contain a")
	}
	ss.Add("a")
	if ss.Size() != 1 {
		t.Fatal("ss size is not 1")
	}
	if ss.Contains("b") {
		t.Fatal("ss contains b")
	}
	ss.Add("b")
	if !ss.Contains("b") {
		t.Fatal("ss does not contain b")
	}
	if ss.Size() != 2 {
		t.Fatal("ss Size is not 2")
	}
	s := ss.ToSlice()
	if len(s) != 2 {
		t.Fatal("s is not length 2: %v", s)
	}
	t.Log("%v %v", s[0], s[1])
	if (s[0] != "a" || s[1] != "b") && (s[0] != "b" || s[1] != "a") {
		t.Fatal("a and b aren't both in the slice: %v", s)
	}
	ss.Remove("a")
	if ss.Contains("a") {
		t.Fatal("ss contains a")
	}
	if ss.Size() != 1 {
 		t.Fatal("ss Size is not 1 after remove")
	}
}
