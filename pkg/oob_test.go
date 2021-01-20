package oob

import (
	"testing"
)

func TestOobPangramLower(t *testing.T) {
	want := "thoob qooboobck broobwn foobx joobmps oobvoobr thoob loobzoob doobg"
	got := Oob("the quick brown fox jumps over the lazy dog")
	if got != want {
		t.Errorf("got \"%s\"; want \"%s\"", got, want)
	}
}

func TestOobPangramUpper(t *testing.T) {
	want := "THOOB QOOBOOBCK BROOBWN FOOBX JOOBMPS OOBVOOBR THOOB LOOBZOOB DOOBG"
	got := Oob("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG")
	if got != want {
		t.Errorf("got \"%s\"; want \"%s\"", got, want)
	}
}
