package slices

import (
	"testing"

	"github.com/12yanogden/strslices"
)

func TestRemoveBeginning(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	expected := []string{"b", "c", "d"}
	actual := Remove(slice, 0)

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestRemoveMiddle(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	expected := []string{"a", "b", "d"}
	actual := Remove(slice, 2)

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestRemoveEnd(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	expected := []string{"a", "b", "c"}
	actual := Remove(slice, (len(slice) - 1))

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}
