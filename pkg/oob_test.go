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

func TestOobPangramMixedCase(t *testing.T) {
	want := "thOOB qoobOOBck broobwn foobx jOOBmps oobvoobr thoob loobzoob dOOBg"
	got := Oob("thE quIck brown fox jUmps over the lazy dOg")
	if got != want {
		t.Errorf("got \"%s\"; want \"%s\"", got, want)
	}
}

// Make sure that we don't double-change "o" instances since "oob" contains "o"
func TestLetterOOrdering(t *testing.T) {
	want := "boobg pooboobts"
	got := Oob("big poots")
	if got != want {
		t.Errorf("got \"%s\"; want \"%s\"", got, want)
	}
}
